package schema

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
"collection1": {
	"_id": "id",
	"create_time": "create_time",
	"update_time": "update_time"
}
*/

var (
	Collection1 = bson.M{
		"name":  "collection1",
		"index": []mongo.IndexModel{},
	}
)
