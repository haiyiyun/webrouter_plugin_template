package serve

import (
	"context"
	"flag"

	"github.com/haiyiyun/webrouter_plugin_template/database/schema"
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

	serveCache := cache.New(serveConf.CacheDefaultExpiration.Duration, serveConf.CacheCleanupInterval.Duration)
	serveDB := mongodb.NewMongoPool("", serveConf.MongoDatabaseName, 100, options.Client().ApplyURI(serveConf.MongoDNS))
	webrouter.SetCloser(func() { serveDB.Disconnect(context.TODO()) })

	serveDB.M().InitCollection(schema.Collection1)
	serveService := serve.NewService(&serveConf, serveCache, serveDB)

	if serveConf.WebRouter {
		//Init Begin
		serveeService1Service := serveService1.NewService(serveService)
		//Init End

		//Go Begin
		//Go End

		//Register Begin
		webrouter.Register(serveConf.WebRouterRootPath+"", serveeService1Service)
		//Register End
	}
}
