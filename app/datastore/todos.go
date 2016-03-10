package datastore

import "gopkg.in/mgo.v2/bson"

func (ds *DataStore) CreateTodo(todo *Todo) error {
	todo.Id = bson.NewObjectId()
	e := ds.DB().C("todos").Insert(todo)
	if e != nil {
		return DSError{e}
	}

	return nil
}

func (ds *DataStore) FindTodo(id bson.ObjectId) (todo *Todo, err error) {
	todo = &Todo{}
	e := ds.DB().C("todos").FindId(id).One(todo)
	if e != nil {
		err = DSError{e}
		return
	}

	return
}

func (ds *DataStore) ListTodos() (todos []Todo, err error) {
	todos = make([]Todo, 0)
	e := ds.DB().C("todos").Find(bson.M{}).All(&todos)
	if e != nil {
		err = DSError{e}
		return
	}

	return
}

func (ds *DataStore) DeleteTodo(id bson.ObjectId) error {
	e := ds.DB().C("todos").RemoveId(id)
	if e != nil {
		return DSError{e}
	}

	return nil
}

func (ds *DataStore) UpdateTodo(todo *Todo) error {
	e := ds.DB().C("todos").UpdateId(todo.Id, todo)
	if e != nil {
		return DSError{e}
	}

	return nil
}
