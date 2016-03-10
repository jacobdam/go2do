package datastore

import "gopkg.in/mgo.v2/bson"

// type ObjectId bson.ObjectId
//
// func (id ObjectId) GetBSON() (interface{}, error) {
// 	return bson.ObjectId(id), nil
// }
//
// func (id *ObjectId) SetBSON(raw bson.Raw) error {
// 	bid := bson.ObjectId("")
// 	if e := raw.Unmarshal(&bid); e != nil {
// 		return e
// 	}
//
// 	*id = ObjectId(bid)
// 	return nil
// }
//
// func (id ObjectId) MarshalJSON() ([]byte, error) {
// 	return bson.ObjectId(id).MarshalJSON()
// }
//
// func (id *ObjectId) UnmarshalJSON(data []byte) error {
// 	bid := bson.ObjectId("")
// 	if e := bid.UnmarshalJSON(data); e != nil {
// 		log.Println(e)
// 		return e
// 	}
//
// 	*id = ObjectId(bid)
// 	return nil
// }

type Token struct {
	Token     string
	UserId    bson.ObjectId
	IsExpired bool
}

// Todo represents a todo item
type Todo struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	ListId      bson.ObjectId `json:"list_id" bson:"list_id"`
	Title       string        `json:"title"`
	IsCompleted bool          `json:"is_completed" bson:"is_completed"`
}

// List represents a list of todo
type List struct {
	Id         bson.ObjectId
	OwnerId    bson.ObjectId
	Title      string
	IsArchived bool
}

// User represents a user in system
type User struct {
	Id             bson.ObjectId
	Username       string
	PasswordDigest string
}

type UserRegistration struct {
	Username string
	Password string
}
