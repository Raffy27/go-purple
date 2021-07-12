package models

import (
	"time"

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
		"email":     user.Email,
		"createdAt": user.CreatedAt,
	})
	secretKey := "haha"
	tokenString, err := token.SignedString([]byte(secretKey))
	return tokenString, err
}
