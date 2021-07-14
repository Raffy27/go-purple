package models

import (
	"context"
	"errors"
	"time"

	"github.com/Raffy27/go-purple/config"
	"github.com/Raffy27/go-purple/server/db"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserExists    = errors.New("user already exists")
	ErrUserNotFound  = errors.New("user not found")
	ErrWrongPassword = errors.New("incorrect password")
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string
	Password string

	Email     string
	CreatedAt time.Time
}

// GetJwtToken returns a JWT token for the given user.
// The returned token contains a username, email, account creation date, ...
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

func FindUser(username string) (*User, error) {
	res := db.C("users").FindOne(context.TODO(), bson.M{"username": username})
	if err := res.Err(); err != nil {
		return nil, err
	}

	var user User
	err := res.Decode(&user)
	return &user, err
}

func FindUserByID(id string) (*User, error) {
	res := db.C("users").FindOne(context.TODO(), bson.M{"_id": id})
	if err := res.Err(); err != nil {
		return nil, err
	}

	var user User
	err := res.Decode(&user)
	return &user, err
}

// Login takes a username and password, computes the hash and attepts to find the given user.
// It returns ErrUserNotFound if no user with the specified username exists.
// It returns ErrWrongPassword if the user does exist, but the hashes differ.
func Login(username, password string) (*User, error) {
	user, err := FindUser(username)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, ErrWrongPassword
	}

	return user, nil
}

// CreateUser creates a new user and returns the ObjectID of the associated document.
func CreateUser(username, password, email string) (primitive.ObjectID, error) {
	// Check if a user with this username already exists
	if _, err := FindUser(username); err != nil {
		if err != mongo.ErrNoDocuments {
			// Something went wrong
			return primitive.NilObjectID, err
		}
	} else {
		return primitive.NilObjectID, ErrUserExists
	}

	user := &User{
		Username:  username,
		Email:     email,
		CreatedAt: time.Now(),
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return primitive.NilObjectID, err
	}
	user.Password = string(hash)

	res, err := db.C("users").InsertOne(context.TODO(), user)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}
