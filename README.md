

## generate mock class
The file mocks/gdbcTemplate.go  and mapper/rowsmapper.go
are generated file by the following approach  
```bash
cd <project_home>
go get github.com/golang/mock/mockgen
cd $GOPATH/src/github.com/golang/mock/mockgen
go build
cd <project_home>
echo $PATH
cp mockgen.exe <User.Home>\go\bin\   #e.g. User.Home = C:\Users\user\
mockgen -destination mocks/gdbcTemplate.go -package mocks -source template/gdbcTemplate.go
mockgen -destination mocks/rowsmapper.go -package mocks -source mapper/rowsmapper.go

```

## create tag

```bash
    git tag v1.x.x
```

## upload tag to the repository
```bash
    git push origin --tags
```

trouble shooting 
in case you encountered the issue:
$ go build
build .: cannot find module for path .

solution:
replace the go build command like 
$ go build mapper/rowsmapper.go


