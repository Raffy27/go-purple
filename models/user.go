package models

import (
	"time"

	"github.com/Raffy27/go-purple/config"
	"github.com/dgrijalva/jwt-go"
)

type User struct {
	//ID     primitive.ObjectID `bson:"_id"`
	Username string
	Password string

	Email     string
	CreatedAt time.Time
}

func (user *User) GetJwtToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  user.Username,
		"email":     user.Email,
		"createdAt": user.CreatedAt,
	})
	secretKey := config.Get().GetString("secrets.jwt")
	tokenString, err := token.SignedString([]byte(secretKey))
	return tokenString, err
}
