# key-value-store

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

## Directory structure

```
  + your_gopath/
  |
  +--+ src/github.com/moemoe89
  |  |
  |  +--+ key-value-store/
  |     |
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
  |        |     + controller_test.go
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
  |        + docker-compese.yml
  |        
  +--+ bin/
  |  |
  |  +-- ... executable file
  |
  +--+ pkg/
     |
     +-- ... all dependency_library required
```