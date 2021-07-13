package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/Raffy27/go-purple/forms"
	"github.com/Raffy27/go-purple/models"
	"github.com/Raffy27/go-purple/models/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct{}

func (auth *Auth) Login(c *gin.Context) {
	var info forms.LoginForm
	if err := c.ShouldBindJSON(&info); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	res := db.Main().Collection("users").FindOne(c, gin.H{
		"username": info.Username,
	})
	switch res.Err() {
	case nil:
		break
	case mongo.ErrNoDocuments:
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Username or password is incorrect."})
		return
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Database error."})
		return
	}
	var user models.User
	res.Decode(&user)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(info.Password)); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Username or password is incorrect."})
		return
	}

	token, _ := user.GetJwtToken()
	c.JSON(http.StatusOK, gin.H{
		"msg":   "success",
		"token": token,
	})
}

func (auth *Auth) Create(c *gin.Context) {
	var info forms.CreateForm
	if err := c.ShouldBindJSON(&info); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Please input all fields",
		})
		return
	}

	user := models.User{
		Username:  info.Username,
		Email:     info.Email,
		CreatedAt: time.Now(),
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(info.Password), bcrypt.MinCost)
	if err != nil {
		log.Panicln(err)
	}
	user.Password = string(hash)

	log.Println(user)
	_, err = db.Main().Collection("users").InsertOne(c, user)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "success!",
	})

}
