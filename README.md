# Sims tree

## Installation

Install [Go](https://go.dev/dl/) v1.18+.

Install the dependencies.

```sh
go mod download
```

Copy .env.

```sh
cp .env.example .env
```

Fill .env properties with your data.

## Updating database schemas

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
