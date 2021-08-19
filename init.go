package webrouter_plugin_template

import (
	"context"
	"flag"

	"webrouter_plugin_template/database/schema"
	"webrouter_plugin_template/service"
	plugin_nameServiceService1 "webrouter_plugin_template/service/service1"

	"github.com/haiyiyun/cache"
	"github.com/haiyiyun/config"
	"github.com/haiyiyun/mongodb"
	"github.com/haiyiyun/webrouter"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	plugin_nameConfFile := flag.String("config.plugin_name", "../config/plugins/plugin_name/plugin_name.conf", "plugin_name config file")
	var plugin_nameConf service.Config
	config.Files(*plugin_nameConfFile).Load(&plugin_nameConf)

	plugin_nameCache := cache.New(plugin_nameConf.CacheDefaultExpiration.Duration, plugin_nameConf.CacheCleanupInterval.Duration)
	plugin_nameDB := mongodb.NewMongoPool("", plugin_nameConf.MongoDatabaseName, 100, options.Client().ApplyURI(plugin_nameConf.MongoDNS))
	webrouter.SetCloser(func() { plugin_nameDB.Disconnect(context.TODO()) })

	plugin_nameDB.M().InitCollection(schema.Collection1)
	plugin_nameService := service.NewService(&plugin_nameConf, plugin_nameCache, plugin_nameDB)

	//Init Begin
	plugin_nameServiceService1Service := plugin_nameServiceService1.NewService(plugin_nameService)
	//Init End

	//Go Begin
	//Go End

	//Register Begin
	webrouter.Register("/", plugin_nameServiceService1Service)
	//Register End
}
