# readme file

a lightweight golang ORM framework similar to jdbc in java tech stack

currently the module support connecting db in mysql driver and PostgreSQL driver.

## introduction

this is a go package to connect kinds of database using template approach.

## usage

install dependencies

```bash
go get github.com/guoapeng/props v1.1.2
go get github.com/guoapeng/gdbcTemplate

```

prepare config.properties unser OS user home directory

for mysql

```properties

DRIVER_NAME=mysql
USERNAME=username
PASSWORD=your_pass
NETWORK=tcp
SERVER=localhost
PORT=3306
DATABASE=database_name
CHARSET=utf8mb4

```

in case you are using postgresql, prepare config.properties as below

```properties

DRIVER_NAME=pgx
USERNAME=postgres
PASSWORD=your_pass
NETWORK=tcp
SERVER=localhost
PORT=5432
DATABASE=database_name
CHARSET=utf8
```

```go

import (
    "github.com/guoapeng/gdbcTemplate/template"
    "github.com/guoapeng/props"
    "log"
)

// sql for postgresql
const SqlUser001 = `select USER_ID,
    USER_NAME,
    NICKNAME,
    PASSWORD,
    USER_TYPE,
    DESCRIPTION
    from USER
    where USER_NAME=$1`

const SqlUser002 = `select USER_ID,
    USER_NAME,
    NICKNAME,
    PASSWORD,
    USER_TYPE,
    DESCRIPTION
    from USER
    where USER_TYPE=$1`

const SqlUser003 = "insert into USER(USER_NAME, NICKNAME, CREATE_DATE, UPDATE_DATE, USER_TYPE)" +
        "values($1, $2, $3, current_date, current_date, 'admin' );"


func main() {

    var err error

    if AppConfig, err = propsReader.NewFactory("application", "config.properties").New(); err != nil {
        log.Fatal("failed to load mandatory properties")
        panic(err)
    }

    gdbcTemplate = template.New(AppConfig)

    // insert without transaction
    result, err = gdbcTemplate.Update(SqlUser003, username, nickname)

    // or insert with transaction
    tx, err := gdbcTemplate.BeginTx()
    result, err = tx.Update(sqlfile.SqlUser003, username, nickname)

    if err != nil {
        tx.Rollback()
        log.Println("ERROR: failed to create user", err)
        return "", fmt.Errorf("failed to create user")
    }

    // query single row
    user := gdbcTemplate.QueryRow(SqlUser001, userName)
    .Map(UserMapper)
    .ToObject()

    // in somce cases, we need to query data thata not commited through transaction,
    // we can do that like the following
    user := tx.QueryRow(SqlUser001, userName)
        .Map(UserMapper)
        .ToObject()

   // query multiple rows
   users := gdbcTemplate.QueryForArray(SqlUser002)
    .Map(mapper.UserRowsMapper, userType).ToArray()

   // query multiple rows within transaction
   users := tx.QueryForArray(SqlUser002)
    .Map(mapper.UserRowsMapper, userType).ToArray()

    tx.Commit()


}

func UserMapper(row *sql.Row) interface{} {
    user := new(domain.User)
    _ = rows.Scan(&user.UserId,
        &user.UserName,
        &user.Nickname,
        &user.Password,
        &user.UserType,
        &user.Description))
    _ = row.Scan(&domainobj.UserId, &domainobj.MemoTarget)

    return user
}

func UserRowsMapper(rows *sql.Rows) interface{} {
    user := new(domain.User)
    _ = rows.Scan(&user.UserId,
        &user.UserName,
        &user.Nickname,
        &user.Password,
        &user.UserType,
        &user.Description)
    return user
}

```

## development guide

### build the go files

```bash
# user project root
go build datasource/datasource.go
go build mapper/rowsmapper.go
go build template/gdbcTemplate.go

# or 

go build ./...
```

### run test

```bash

go test template/gdbcTemplate_test.go
go test transaction/transaction_test.go
go test datasource/dbManager_test.go

# or

go test ./...

```

### generate mock class

generate mock classes with testify

```bash

# install mockery

cd <project_home>
go install github.com/vektra/mockery/v2@v2.25.0

cd <project_home>
go get github.com/stretchr/testify/mock
go get github.com/vektra/mockery/.../

cd <project_home>
# generate mock structs
mockery --recursive --name "GdbcTemplate|RowMapper|RowsMapper|Transaction|DataSource|ConnManager"

mockery -r -all

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

## todo list

1. remove dependency on github.com/guoapeng/props to only rely on sql.datasource
2. support sqlite

## trouble shooting

1. cannot find module
   in case you encountered the issue:

```bash
   $ go build
   build .: cannot find module for path .
```

solution:
replace the go build command like

```bash
go build mapper/rowsmapper.go
```
