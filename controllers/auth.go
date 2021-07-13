package controllers

import (
	"log"
	"net/http"

	"github.com/Raffy27/go-purple/forms"
	"github.com/Raffy27/go-purple/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Auth struct{}

func (auth *Auth) Login(c *gin.Context) {
	// Validate form input
	var info forms.LoginForm
	if err := c.ShouldBindJSON(&info); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	// Attempt login
	svc := &models.UserMethods{Context: c}
	user, err := svc.Login(info.Username, info.Password)
	log.Println(err)
	switch err {
	case nil:
		break
	case mongo.ErrNoDocuments:
		// User not found
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Username or password is incorrect."})
		return
	default:
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Derive token
	token, err := user.GetJwtToken()
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error deriving token."})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":   "success",
		"token": token,
	})
}

func (auth *Auth) Create(c *gin.Context) {
	// Validate form input
	var info forms.CreateForm
	if err := c.ShouldBindJSON(&info); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Please input all fields",
		})
		return
	}

	// Attempt to create user
	svc := &models.UserMethods{Context: c}
	id, err := svc.Create(info.Username, info.Password, info.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Created user %s with name '%s'", id, info.Username)

	// Success
	c.JSON(http.StatusOK, gin.H{
		"msg": "success!",
	})

}
