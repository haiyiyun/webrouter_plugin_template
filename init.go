package webrouter_plugin_template

import (
	"context"
	"flag"

	"github.com/haiyiyun/webrouter_plugin_template/database/schema"
	"github.com/haiyiyun/webrouter_plugin_template/service"
	webrouter_plugin_templateServiceService1 "github.com/haiyiyun/webrouter_plugin_template/service/service1"

	"github.com/haiyiyun/cache"
	"github.com/haiyiyun/config"
	"github.com/haiyiyun/mongodb"
	"github.com/haiyiyun/webrouter"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	webrouter_plugin_templateConfFile := flag.String("config.webrouter_plugin_template", "../config/plugins/webrouter_plugin_template/webrouter_plugin_template.conf", "webrouter_plugin_template config file")
	var webrouter_plugin_templateConf service.Config
	config.Files(*webrouter_plugin_templateConfFile).Load(&webrouter_plugin_templateConf)

	webrouter_plugin_templateCache := cache.New(webrouter_plugin_templateConf.CacheDefaultExpiration.Duration, webrouter_plugin_templateConf.CacheCleanupInterval.Duration)
	webrouter_plugin_templateDB := mongodb.NewMongoPool("", webrouter_plugin_templateConf.MongoDatabaseName, 100, options.Client().ApplyURI(webrouter_plugin_templateConf.MongoDNS))
	webrouter.SetCloser(func() { webrouter_plugin_templateDB.Disconnect(context.TODO()) })

	webrouter_plugin_templateDB.M().InitCollection(schema.Collection1)
	webrouter_plugin_templateService := service.NewService(&webrouter_plugin_templateConf, webrouter_plugin_templateCache, webrouter_plugin_templateDB)

	//Init Begin
	webrouter_plugin_templateServiceService1Service := webrouter_plugin_templateServiceService1.NewService(webrouter_plugin_templateService)
	//Init End

	//Go Begin
	//Go End

	//Register Begin
	webrouter.Register("/", webrouter_plugin_templateServiceService1Service)
	//Register End
}
