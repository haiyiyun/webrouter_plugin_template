package manage

import (
	"context"
	"flag"

	"github.com/haiyiyun/webrouter_plugin_template/database/schema"
	"github.com/haiyiyun/webrouter_plugin_template/service/manage"
	manageService1 "github.com/haiyiyun/webrouter_plugin_template/service/manage/service1"

	"github.com/haiyiyun/cache"
	"github.com/haiyiyun/config"
	"github.com/haiyiyun/mongodb"
	"github.com/haiyiyun/webrouter"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	manageConfFile := flag.String("config.webrouter_plugin_template.manage", "../config/plugins/webrouter_plugin_template/manage.conf", "manage config file")
	var manageConf manage.Config
	config.Files(*manageConfFile).Load(&manageConf)

	manageCache := cache.New(manageConf.CacheDefaultExpiration.Duration, manageConf.CacheCleanupInterval.Duration)
	manageDB := mongodb.NewMongoPool("", manageConf.MongoDatabaseName, 100, options.Client().ApplyURI(manageConf.MongoDNS))
	webrouter.SetCloser(func() { manageDB.Disconnect(context.TODO()) })

	manageDB.M().InitCollection(schema.Collection1)
	manageService := manage.NewService(&manageConf, manageCache, manageDB)

	if manageConf.WebRouter {
		//Init Begin
		manageeService1Service := manageService1.NewService(manageService)
		//Init End

		//Go Begin
		//Go End

		//Register Begin
		webrouter.Register(manageConf.WebRouterRootPath+"", manageeService1Service)
		//Register End
	}
}
