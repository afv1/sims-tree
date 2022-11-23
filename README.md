# Sims tree

## Installation

Install [Go](https://go.dev/dl/) v1.18+.

Install the dependencies.

```sh
go mod download
```

Copy .env file.

```sh
cp .env.example .env
```

Create Docker images for DB and Redis (working on complete Docker container)

Set available ports for images! (e.g. postgre: -p xxxx:5432, redis: -p xxxx:6379)
```sh
docker run --name=sims-tree-db -e POSTGRES_PASSWORD='root' -p 5433:5432 -d --rm postgres
docker run -d -p 6381:6379 redislabs/redismod
```

Fill .env properties with your data.

## DB schema

See [Ent.](https://entgo.io/docs/getting-started/) docs.

## Serve

```sh
go run cmd/main.go
```

## Building

```sh
go build -o sims-tree
```

## Run built Server App

#### Linux
```sh
./sims-tree
```

#### Shindows
```sh
run sims-tree.exe
```

## Install Server App

```sh
go install
```

## Run installed Server App

#### Linux
```sh
sims-tree
```
