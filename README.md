# readme file

currently the module support connecting db in mysql driver and PostgreSQL driver.

## introduction

this is a go package to connect kinds of database using template approach.

## usage

install dependencies

```bash
go get github.com/guoapeng/props v1.1.2
go get github.com/guoapeng/gdbcTemplate

```

```go

import (
    "github.com/guoapeng/gdbcTemplate/template"
    "github.com/guoapeng/props"
    "log"
)


func main() {
    var err error

    if AppConfig, err = propsReader.NewFactory("application", "config.properties").New(); err != nil {
        log.Fatal("failed to load mandatory properties")
        panic(err)
    }

    gdbcTemplate = template.New(AppConfig)
    gdbcTemplate.QueryRow("select * from USERS where USER_NAME=?", "your_name")
    .Map(func(rows *sql.Row) interface{} { return "test" })
    .ToObject()

    gdbcTemplate.QueryForArray("select * from TABLE_NAME where USER_NAME=?", "your_name")
    .Map(func(rows *sql.Rows) interface{} { return "test" }).ToArray()

}
```

## development guide

### build the go files

```bash
# user project root
go build datasource/datasource.go
go build mapper/rowsmapper.go
go build template/gdbcTemplate.go
```

### run test

```bash

go test template/gdbcTemplate_test.go

```

### generate mock class

The file mocks/gdbcTemplate.go  and mapper/rowsmapper.go
are generated file by the following approach  

```bash
cd <project_home>
go get github.com/golang/mock/mockgen
go install github.com/golang/mock/mockgen
cd <project_home>

mockgen -destination mocks/gdbcTemplate.go -package mocks -source template/gdbcTemplate.go
mockgen -destination mocks/rowsmapper.go -package mocks -source mapper/rowsmapper.go
$mockgen -destination mocks/appConfig.go -package mocks github.com/guoapeng/props AppConfigProperties

```

## publish

### create tag

```bash
    git tag v1.x.x
```

### upload tag to the repository

```bash
    git push origin --tags
```

## trouble shooting

1. cannot find module
in case you encountered the issue:
$ go build
build .: cannot find module for path .

solution:
replace the go build command like
$ go build mapper/rowsmapper.go
