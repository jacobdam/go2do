package datastore

import "gopkg.in/mgo.v2/bson"

type ObjectId bson.ObjectId

type Token struct {
	token     string
	UserId    ObjectId
	IsExpired bool
}

type Todo struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	ListId      ObjectId      `json:"list_id" binding:"required"`
	Title       string        `json:"title"`
	IsCompleted bool          `json:"is_completed"`
}

type TodoList struct {
	Id         ObjectId
	OwnerId    ObjectId
	Title      string
	IsArchived bool
}

type User struct {
	Id             ObjectId
	Username       string
	PasswordDigest string
}

type UserRegistration struct {
	Username string
	Password string
}
