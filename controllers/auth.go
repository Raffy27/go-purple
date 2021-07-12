package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/Raffy27/go-purple/models"
	"github.com/Raffy27/go-purple/models/db"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct{}

func (auth *Auth) Login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	res := db.Main().Collection("users").FindOne(c, gin.H{})
	if err := res.Err(); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err,
		})
		return
	}

	token, _ := user.GetJwtToken()
	c.JSON(http.StatusOK, gin.H{
		"msg":   "success",
		"token": token,
	})
}

func (auth *Auth) Create(c *gin.Context) {
	type createInfo struct {
		User string `binding:"required"`
		Pass string `binding:"required"`

		Email string `binding:"required"`
	}
	var info createInfo

	if err := c.ShouldBindJSON(&info); err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"error": "Please input all fields",
		})
		return
	}

	user := models.User{
		Username:  info.User,
		Email:     info.Email,
		CreatedAt: time.Now(),
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(info.Pass), bcrypt.MinCost)
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
