package datastore

type ObjectId string

type Token struct {
	token     string
	UserId    ObjectId
	IsExpired bool
}

type Todo struct {
	Id          string
	ListId      ObjectId
	Title       string
	IsCompleted bool
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
