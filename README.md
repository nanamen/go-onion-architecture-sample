### go-echo-rest-sample
go-echo-rest-sample(Go, echo, REST, DDD)

## Depends on
* echo
* viper
* gorm
* goose
* assert
* go-sqlmock.v1

## How to use
### Install
* Dep  
`$ brew install dep`  
* go-task  
`$ go get -u -v github.com/go-task/task/cmd/task`  
* gometalinter  
`$ go get -u -v github.com/alecthomas/gometalinter`  
* goose  
`$ go get bitbucket.org/liamstask/goose/cmd/goose`  

### Application setting
`$ gometalinter --install --update`  
`$ dep ensure`

### Database setting   
You have to write Database connection information in db/dbconf.yml and conf/local.yml  
`$ goose up`

### Run Server
```
task run
```

## Sample Request
##### Create User
```
curl -X POST \
  -H 'Content-Type: application/json' \
  -d '{"name":"Joe Smith"}' \
  localhost:1323/users
```

##### Get AllUser

`curl localhost:1323/users`

##### Get User

`curl localhost:1323/users/1`


##### Update User
```
curl -X PUT \
  -H 'Content-Type: application/json' \
  -d '{"name":"Joe"}' \
  localhost:1323/users/1
```

##### Delete User
`curl -X DELETE localhost:1323/users/1`