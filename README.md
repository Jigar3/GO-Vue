# Vue & Go

## This is a simple Blogging Application made using a GO server as backend and Vue as the frontend framework. The database of choice was MongoDB.

## Backend

~~Dependency Management in Go is not the best and I didn't get around using Go Dep and just put the whole folder inside **$GOPATH/src/github.com/\<myusername>/\<this folder containing backend & frontend>**~~

Now it uses Go Dep so no problems. Clone it anywhere you want.

Go dependencies used are Gin, Cors, Mgo, bson and standard packages like time, json & log


You can run the go server using 
```
cd backend/
go run main.go routes.go
```

Store your Mongo URI in an environment variable MONGO_DB_URL
For Bash
```
export MONGO_DB_URL="mongodb://localhost:27017/<dbname>"
```

## Frontend

Frontend is made using Vue.js & to run the the Vue server use the below commands

```
cd frontend/
yarn serve
```
