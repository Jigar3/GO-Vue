package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"time"
)

// Blog struct
type Blog struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Title     string        `bson:"title" json:"title" binding:"required"`
	Content   string        `bson:"content" json:"content" binding:"required"`
	CreatedAt string        `bson:"time" json:"time"`
	Author    string        `bson:"author" json:"author"`
}

// DB export
var DB *mgo.Database

// NewRouter export
func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", sayAlive)
	router.GET("/view", getAll)
	router.GET("/view/:id", getParticular)
	router.POST("/new", createPost)
	router.DELETE("/delete/:id", deletePost)
	router.PATCH("/update/:id", updatePost)
	router.DELETE("/eject", destruct)

	return router
}

func sayAlive(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "I am alive",
	})
}

func createPost(c *gin.Context) {
	var blog Blog

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Missing request parameters",
		})
		return
	}

	blog.ID = bson.NewObjectId()
	blog.CreatedAt = time.Now().String()
	// blog.Author = "Jigar Chavada"

	if err := DB.C("posts").Insert(blog); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Something went wrong",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Inserted in DB",
	})
}

func getAll(c *gin.Context) {
	var blog []Blog

	if err := DB.C("posts").Find(nil).All(&blog); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No Data Found",
		})
	}

	c.JSON(http.StatusOK, blog)
}

func getParticular(c *gin.Context) {
	id := bson.ObjectIdHex(c.Param("id"))
	blog := Blog{}
	// log.Printf(id)

	if err := DB.C("posts").Find(bson.M{"_id": id}).One(&blog); err != nil {
		log.Print(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No blog found",
		})

		return
	}

	c.JSON(http.StatusOK, blog)
}

func deletePost(c *gin.Context) {
	id := bson.ObjectIdHex(c.Param("id"))

	if err := DB.C("posts").Remove(bson.M{"_id": id}); err != nil {
		log.Print(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Couldn't find the blog",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted Sucessfully",
	})
}

func updatePost(c *gin.Context) {
	type blogUpdate struct {
		Title   string `bson:"title json:"title" binding:"required"`
		Content string `bson:"content" json:"content" binding:"required"`
	}

	var blog blogUpdate

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect Parameters",
		})

		return
	}

	log.Print(blog.Content)

	id := bson.ObjectIdHex(c.Param("id"))

	if err := DB.C("posts").Update(bson.M{"_id": id}, bson.M{"$set": bson.M{"content": blog.Content, "title": blog.Title}}); err != nil {
		log.Print(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Can't find the blog to update",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Updated Succesfully",
	})
}

func destruct(c *gin.Context) {

	if _, err := DB.C("posts").RemoveAll(nil); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Unsuccessful",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successful",
	})
}
