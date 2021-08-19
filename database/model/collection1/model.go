package collection1

import (
	"webrouter_plugin_template/database"
	"webrouter_plugin_template/database/schema"

	"github.com/haiyiyun/mongodb"
)

type Model struct {
	*database.Database `json:"-" bson:"-" map:"-"`
}

func NewModel(mgo mongodb.Mongoer) *Model {
	obj := &Model{
		Database: database.NewDatabase(mgo, schema.Collection1),
	}

	return obj
}
