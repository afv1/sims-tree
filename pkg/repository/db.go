package repository

import (
    "context"
    "fmt"
    "github.com/afv1/sims-tree/cfg"
    "github.com/afv1/sims-tree/ent"
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/lib/pq"
    _ "github.com/mattn/go-sqlite3"
    "github.com/sirupsen/logrus"
)

func buildDsn() (string, error) {
    var dsn string

    drvr := cfg.App.DBConfig.Driver
    user := cfg.App.DBConfig.User
    pswd := cfg.App.DBConfig.Password
    host := cfg.App.DBConfig.Host
    port := cfg.App.DBConfig.Port
    name := cfg.App.DBConfig.Name
    sslm := cfg.App.DBConfig.SSLMode

    if drvr == "mysql" {
        frmt := "%s:%s@tcp(%s:%s)/%s?parseTime=True"
		dsn = fmt.Sprintf(frmt, user, pswd, host, port, name)
	} else if drvr == "postgres" {
        frmt := "host=%s port=%s user=%s dbname=%s password=%s sslmode=%s"
		dsn = fmt.Sprintf(frmt, host, port, user, name, pswd, sslm)
	} else if drvr == "sqlite" {
		dsn = "file:ent?mode=memory&cache=shared&_fk=1"
	} else {
		return "", fmt.Errorf("driver %s is not supported", drvr)
	}

    return dsn, nil
}

func GetClient() *ent.Client {
    dsn, err := buildDsn()

    if err != nil && dsn == "" {
        logrus.Fatalf("%s", err.Error())
    }

    client, err := ent.Open(cfg.App.DBConfig.Driver, dsn)

    if err != nil {
        logrus.Fatalf("Couild not connect to DB: %s", err.Error())
    }

    return client
}

func CloseClient(client *ent.Client) {
    e := client.Close()

    if e != nil {
        logrus.Fatalf("Could not close Ent client: %s", e.Error())
    }
}

func Migrate() {
    client := GetClient()
    defer CloseClient(client)

    if err := client.Schema.Create(context.Background()); err != nil {
        logrus.Fatalf("Failed creating schema resources: %v", err)
    }
}
