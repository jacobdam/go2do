package datastore

import (
	"log"

	"gopkg.in/mgo.v2/bson"
)

func (ds *DataStore) CreateTodo(todo *Todo) (err error) {
	// todo.Id = ObjectId(bson.NewObjectId())
	// if e := ds.DB().C("todos").Insert(todo); e != nil {
	i := bson.NewObjectId()
	doc := bson.M{"_id": i, "foo": "bar"}
	if e := ds.DB().C("todos").Insert(doc); e != nil {
		err = DSError{e}
		return
	}

	log.Println(todo)

	return
}

// func (ds *DataStore) FindTodo(id ObjectId) (todo *TODO, err error) {
//
// }
//
// func (ds *DataStore) DeleteTodo(id ObjectId) (err error) {
//
// }
//
// func (ds *DataStore) UpdateTodo(id ObjectId, todo *Todo) (err error) {
//
// }
