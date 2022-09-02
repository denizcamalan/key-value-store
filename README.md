# key-value-store

This is an in-memory, key/value store application. Think of it as a dictionary with any number of keys, each of which has a value that can be set or retrieved. Used Redis as a store. 

The application will use HTTP status codes to communicate the success or failure of an operation.

## Getting Started

1. git clone https://github.com/denizcamalan/key-value-store.git

2. cd  key-value-store

3. Run application
```
docker-compese up -d
```
4.  Browse Swagger UI [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

5. Run Tests
```
go test ./...
```
![Screen Shot 2022-08-26 at 12 43 17 PM](https://user-images.githubusercontent.com/79871039/186877246-238ddaf8-4482-43b7-b767-472aa5d69d4f.png)

# Technologies in use

  Go Programming Language, Redis, Docker, Swagger Documantation

## Directory structure

```

  |  +--+ key-value-store/
  |     |
  |     + docker-compese.yml
        |
  |     +--+ src
  |        +--+ main.go
  |        +--+ configuration
  |        |  |
  |        |  +--+ cors_conf.go
  |        |     + redis_conf_test.go
  |        |     + redis_conf.go
  |        |     + viper_conf.go
  |        +--+ controller
  |        |  |
  |        |  +--+ controller.go
  |        +--+ docs
  |        |  |
  |        |  +--+ docs.go
  |        |     + swagger.json
  |        |     + swagger.yaml
  |        +--+ model
  |        |  |
  |        |  +--+ message.go
  |        |     + workflow.go
  |        +--+ repository
  |        |  |
  |        |  +--+ respitory.go
  |        |     + respitory_test.go
  |        |
  |        +--+ router
  |           |
  |           +--+ router.go
  |        +--+ resource
  |           |
  |           +--+ properties-prod.yaml
  |              + properties-prod-staging.yaml

```
