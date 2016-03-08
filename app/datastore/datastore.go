package datastore

import (
	"strings"

	"gopkg.in/mgo.v2"
)

type DSConfig struct {
	urls []string
}

type DataStore struct {
	session *mgo.Session
	config  *DSConfig
}

type DSError struct {
	error
}

func NewDSConfig(urls ...string) DSConfig {
	return DSConfig{urls: urls}
}

func NewDataStore(cfg *DSConfig) (ds *DataStore, err error) {
	session, e := mgo.Dial(strings.Join(cfg.urls, ","))
	if e != nil {
		err = DSError{e}
		return
	}

	ds = &DataStore{session: session, config: cfg}
	return
}

func (ds *DataStore) Copy() *DataStore {
	return &DataStore{ds.session.Copy(), ds.config}
}

func (ds *DataStore) Close() {
	ds.session.Close()
}

func (ds *DataStore) DB() *mgo.Database {
	return ds.session.DB("")
}
