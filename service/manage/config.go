package manage

import (
	"github.com/haiyiyun/webrouter_plugin_template/service/base"
)

type WebrouterPluginTemplateConfig struct {
	WebRouter         bool   `json:"web_router"`
	WebRouterRootPath string `json:"web_router_root_path"`
}

type Config struct {
	base.Config
	WebrouterPluginTemplateConfig
}
