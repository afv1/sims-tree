package main

import (
    simstree "github.com/afv1/sims-tree"
    "github.com/afv1/sims-tree/cfg"
    "github.com/afv1/sims-tree/pkg/handler"
    "github.com/afv1/sims-tree/pkg/repository"
    "github.com/afv1/sims-tree/pkg/service"
    "github.com/joho/godotenv"
    "github.com/sirupsen/logrus"
)

func main() {
    // set logger
    logrus.SetFormatter(new(logrus.JSONFormatter))

    err := initApp()

    if err != nil {
        logrus.Fatalf("Could not init app: %s", err.Error())
    }

    server := new(simstree.Server)
    services := service.NewProvider()
    handlers := handler.NewHandler(services)

    if err = server.Run(handlers.InitRoutes()); err != nil {
        logrus.Fatalf("Could not run server: %s", err.Error())
    }
}

// initApp load .env, init configs and migrate DB Schema
func initApp() error {
    err := godotenv.Load()

    if err != nil {
        return err
    }

    cfg.InitApp()
    repository.Migrate()

    return nil
}