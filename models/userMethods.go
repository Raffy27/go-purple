package models

import (
	"context"
	"errors"
	"time"

	"github.com/Raffy27/go-purple/models/db"
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

type UserMethods struct {
	context.Context
}

func (c *UserMethods) FindByUsername(username string) (*User, error) {
	res := db.Main().Collection("users").FindOne(c, bson.M{"username": username})
	if err := res.Err(); err != nil {
		return nil, err
	}

	var user User
	err := res.Decode(&user)
	return &user, err
}

func (c *UserMethods) FindByID(id string) (*User, error) {
	res := db.Main().Collection("users").FindOne(c, bson.M{"_id": id})
	if err := res.Err(); err != nil {
		return nil, err
	}

	var user User
	err := res.Decode(&user)
	return &user, err
}

func (c *UserMethods) Login(username, password string) (*User, error) {
	user, err := c.FindByUsername(username)
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

// Create creates a new user and returns the ObjectID of the associated document
func (c *UserMethods) Create(username, password, email string) (primitive.ObjectID, error) {
	// Check if a user with this username already exists
	if _, err := c.FindByUsername(username); err != nil {
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

	res, err := db.Main().Collection("users").InsertOne(c, user)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}
