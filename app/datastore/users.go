package datastore

import "golang.org/x/crypto/bcrypt"

func (ds *DataStore) CreateUser(ur *UserRegistration) *User {
	pwDigest, _ := bcrypt.GenerateFromPassword([]byte(ur.Password), 0)
	user := User{Username: ur.Username, PasswordDigest: string(pwDigest)}
	ds.DB().C("users").Insert(user)

	return &user
}
