package config

import (
	"os"

	"github.com/jacobdam/go2do/core"
)

var DB = map[string]core.DSConfig{
	"development": core.NewDSConfig("localhost/go2do"),
	"test":        core.NewDSConfig("localhost/go2do_test"),
	"production":  core.NewDSConfig(os.Getenv("MONGODB_URI")),
}
