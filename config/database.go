package config

import (
	"os"

	"github.com/jacobdam/go2do/app/datastore"
)

var DB = map[string]datastore.DSConfig{
	"development": datastore.NewDSConfig("localhost/go2do"),
	"test":        datastore.NewDSConfig("localhost/go2do_test"),
	"production":  datastore.NewDSConfig(os.Getenv("MONGODB_URI")),
}
