package serve

import (
	"context"
	"flag"

	"github.com/haiyiyun/webrouter_plugin_template/database/schema"
	"github.com/haiyiyun/webrouter_plugin_template/service/base"
	"github.com/haiyiyun/webrouter_plugin_template/service/serve"
	serveService1 "github.com/haiyiyun/webrouter_plugin_template/service/serve/service1"

	"github.com/haiyiyun/cache"
	"github.com/haiyiyun/config"
	"github.com/haiyiyun/mongodb"
	"github.com/haiyiyun/webrouter"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	serveConfFile := flag.String("config.webrouter_plugin_template.serve", "../config/plugins/webrouter_plugin_template/serve.conf", "serve config file")
	var serveConf serve.Config
	config.Files(*serveConfFile).Load(&serveConf)

	if serveConf.WebRouter {
		baseConfFile := flag.String("config.plugins.webrouter_plugin_template.serve.base", "../config/plugins/webrouter_plugin_template/base.conf", "base config file")
		var baseConf base.Config
		config.Files(*baseConfFile).Load(&baseConf)

		baseCache := cache.New(baseConf.CacheDefaultExpiration.Duration, baseConf.CacheCleanupInterval.Duration)
		baseDB := mongodb.NewMongoPool("", baseConf.MongoDatabaseName, 100, options.Client().ApplyURI(baseConf.MongoDNS))
		webrouter.SetCloser(func() { baseDB.Disconnect(context.TODO()) })

		baseDB.M().InitCollection(schema.Collection1)

		baseService := base.NewService(&baseConf, baseCache, baseDB)

		serveConf.Config = baseConf
		serveService := serve.NewService(&serveConf, baseService)

		//Init Begin
		serveService1Service := serveService1.NewService(serveService)
		//Init End

		//Go Begin
		//Go End

		//Register Begin
		webrouter.Register(serveConf.WebRouterRootPath+"", serveService1Service)
		//Register End
	}
}
