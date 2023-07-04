# Implementación de SEArch

## Componentes

Por ahora habrá estos componentes en la arquitectura:

- broker + repository
- middleware

## Código generado (go protobufs y go-grpc)

### Buf para protobufs y gRPC

    buf generate proto

### Entgo para manejo de la base de datos del broker

    go generate ./ent

### Para regenerar mocks con [mockery](https://vektra.github.io/mockery/)

    mockery --dir contract --all --with-expecter

## Comandos varios

### To get a report of code coverage

    go test ./... -coverprofile=coverage.txt -covermode atomic -race -timeout 4s
    go tool cover -html=coverage.txt

### Para correr los tests

    go test ./...

Y con el [race detector](https://go.dev/doc/articles/race_detector):

    go test ./... -count=1 -race

### Para compilar los binarios de broker y middleware

    go build ./cmd/broker ./cmd/middleware

### Comandos útiles de Entgo (ORM)

#### Show schema in CLI

    go run -mod=mod entgo.io/ent/cmd/ent describe ./ent/schema

#### Show schema in [Atlas Cloud](https://gh.atlasgo.cloud/)

    go run -mod=mod ariga.io/entviz ./ent/schema

#### Generate Entity Relation diagram locally

    go run -mod=mod github.com/a8m/enter ./ent/schema

