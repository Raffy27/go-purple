package controllers

import (
	"github.com/Raffy27/go-purple/models"
	"github.com/Raffy27/go-purple/server/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserController struct{}

func (*UserController) GetAll(c *gin.Context) {
	opt := &options.FindOptions{Projection: bson.M{
		"password": 0,
	}}
	cursor, err := db.C("users").Find(c, bson.M{}, opt)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(c)

	var users []models.User
	err = cursor.All(c, &users)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"users": users})
}

func (*UserController) GetByUsername(c *gin.Context) {
	name := c.Param("user")
	user, err := models.FindUser(name)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	// Remove unncessary fields
	user.Password = ""

	c.JSON(200, gin.H{"user": user})
}
