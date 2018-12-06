# Vue & Go

## This is a simple Blogging Application made using a GO server as backend and Vue as the frontend framework. The database of choice was MongoDB.

## Backend

Dependency Management in Go is not the best and I didn't get around using Go Dep and just put the whole folder inside **$GOPATH/src/github.com/\<myusername>/\<this folder containing backend & frontend>**

Go dependencies used are Gin, Cors, Mgo, bson and standard packages like time, json & log

```
go get "github.com/gin-contrib/cors"
go get "github.com/gin-gonic/gin"
go get "gopkg.in/mgo.v2"
go get "gopkg.in/mgo.v2/bson"
```

You can run the go server using 
```
go run main.go routes.go
```



## Frontend

Frontend is made using Vue.js & to run the the Vue server use the below commands

```
cd frontend/
yarn serve
```
