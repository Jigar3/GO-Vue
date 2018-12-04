package some

import (
	"fmt"
	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

type Student struct {
	Id         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name       string
	Enrollment int32
}

func main() {

	router := gin.Default()
	router.GET("/insert", InsertIntoDb)
	router.GET("/view", findAll)
	router.Run(":8080")

	// fmt.Println(bson.NewObjectId().Hex())
}

func getSession() *mgo.Collection {
	uri := "mongodb://localhost:27017/go-demo"
	sess, err := mgo.Dial(uri)

	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	// sess.SetSafe(&mgo.Safe{})

	collection := sess.DB("go-demo").C("users")

	return collection
}

// InsertIntoDb into db
func InsertIntoDb(c *gin.Context) {
	uri := "mongodb://localhost:27017/go-demo"
	sess, err := mgo.Dial(uri)

	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	// sess.SetSafe(&mgo.Safe{})

	collection := sess.DB("go-demo").C("users")

	// return collection

	newStudent := Student{
		bson.NewObjectId(),
		"Ekansh",
		15,
	}

	err = collection.Insert(newStudent)

	if err != nil {
		log.Fatal("Problem inserting data: ", err)
		return
	}
	fmt.Println("Mast")

	count, _ := collection.Count()

	c.JSON(http.StatusOK, gin.H{
		"Status": "Inserted",
		"data":   newStudent,
		"count":  count,
	})
}

func findAll(c *gin.Context) {

	collection := getSession()

	var result []Student
	_ = collection.Find(nil).All(&result)

	c.JSON(http.StatusOK, result)
}

// func removeAll(c *gin.Context) {

// }
