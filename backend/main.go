package main

import (
	mgo "gopkg.in/mgo.v2"
	"log"
	"os"
)

func main() {
	router := NewRouter()

	session, err := mgo.Dial(os.Getenv("MONGO_DB_URL"))

	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	DB = session.DB("go-demo")

	router.Run(":8000")

}
