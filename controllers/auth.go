package controllers

import (
	"log"
	"net/http"

	"github.com/Raffy27/go-purple/forms"
	"github.com/Raffy27/go-purple/models"
	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func (*AuthController) Login(c *gin.Context) {
	// Validate form input
	var info forms.LoginForm
	if err := c.ShouldBindJSON(&info); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	// Attempt login
	user, err := models.Login(info.Username, info.Password)
	switch err {
	case nil:
		break
	case models.ErrUserNotFound:
		fallthrough
	case models.ErrWrongPassword:
		c.AbortWithStatusJSON(401, gin.H{"error": "Username or password is incorrect"})
		return
	default:
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	// Derive token
	token, err := user.GetJwtToken()
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(500, gin.H{"error": "Error deriving token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":   "success",
		"token": token,
	})
}

func (*AuthController) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}

func (*AuthController) Register(c *gin.Context) {
	// Validate form input
	var info forms.RegisterForm
	if err := c.ShouldBindJSON(&info); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Please input all fields",
		})
		return
	}

	// Attempt to create user
	id, err := models.CreateUser(info.Username, info.Password, info.Email)
	switch err {
	case nil:
		break
	case models.ErrUserExists:
		c.AbortWithStatusJSON(400, gin.H{"error": "Username already exists"})
		return
	default:
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Created user [%s] with name '%s'", id.Hex(), info.Username)

	// Success
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})

}
