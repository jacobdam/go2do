package datastore

import "golang.org/x/crypto/bcrypt"

func genPwDigest(pw string) string {
	b, _ := bcrypt.GenerateFromPassword([]byte(pw), 0)

	return string(b)
}

func (ds *DataStore) CreateUser(ur *UserRegistration) (user *User, err error) {
	u := User{Username: ur.Username, PasswordDigest: genPwDigest(ur.Password)}
	e := ds.DB().C("users").Insert(&u)
	if e != nil {
		err = DSError{e}
		return
	}

	user = &u
	return
}
