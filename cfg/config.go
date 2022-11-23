package cfg

import (
    "os"
    "strconv"
)

var App *app

type app struct {
    Host               string
    Port               string
    ReadTimeout        int
    WriteTimeout       int
    AuthSessTTL        int
    AuthSessKey        string
    DBConfig           DBCfg
    RedisPersistConfig RedisPersistCfg
    RedisCacheConfig   RedisCacheCfg
}

type DBCfg struct {
    Driver       string
    Host         string
    Port         string
    Password     string
    RootPassword string
    User         string
    Name         string
    SSLMode      string
}

type RedisPersistCfg struct {
    Host     string
    Port     string
    Password string
    DB       string
    PoolSize string
}

type RedisCacheCfg struct {
    Host     string
    Port     string
    Password string
    DB       string
    PoolSize string
}

func initDB() DBCfg {
    return DBCfg{
        Driver:   os.Getenv("DB_DRIVER"),
        Host:     os.Getenv("DB_HOST"),
        Port:     os.Getenv("DB_PORT"),
        User:     os.Getenv("DB_USER"),
        Password: os.Getenv("DB_PASS"),
        Name:     os.Getenv("DB_NAME"),
        SSLMode:  os.Getenv("DB_SSL_MODE"),
    }
}

func initRedisPersist() RedisPersistCfg {
    return RedisPersistCfg{
        Host:     os.Getenv("REDIS_PERSIST_HOST"),
        Port:     os.Getenv("REDIS_PERSIST_PORT"),
        Password: os.Getenv("REDIS_PERSIST_PASSWORD"),
        DB:       os.Getenv("REDIS_PERSIST_DB"),
        PoolSize: os.Getenv("REDIS_PERSIST_POOLSIZE"),
    }
}

func initRedisCache() RedisCacheCfg {
    return RedisCacheCfg{
        Host:     os.Getenv("REDIS_CACHE_HOST"),
        Port:     os.Getenv("REDIS_CACHE_PORT"),
        Password: os.Getenv("REDIS_CACHE_PASSWORD"),
        DB:       os.Getenv("REDIS_CACHE_DB"),
        PoolSize: os.Getenv("REDIS_CACHE_POOLSIZE"),
    }
}

func InitApp() {
    cookieTtl, _ := strconv.Atoi(os.Getenv("AUTH_COOKIE_TTL"))
    readTimeout, _ := strconv.Atoi(os.Getenv("APP_READ_TIMEOUT"))
    writeTimeout, _ := strconv.Atoi(os.Getenv("APP_WRITE_TIMEOUT"))

    App = &app{
        Host:               os.Getenv("APP_HOST"),
        Port:               os.Getenv("APP_PORT"),
        ReadTimeout:        readTimeout,
        WriteTimeout:       writeTimeout,
        AuthSessTTL:        cookieTtl,
        AuthSessKey:        os.Getenv("AUTH_COOKIE_KEY"),
        DBConfig:           initDB(),
        RedisPersistConfig: initRedisPersist(),
        RedisCacheConfig:   initRedisCache(),
    }
}