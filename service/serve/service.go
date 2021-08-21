package serve

import "github.com/haiyiyun/webrouter_plugin_template/service/base"

type Service struct {
	*Config
	*base.Service
}

func NewService(c *Config, s *base.Service) *Service {
	return &Service{
		Config:  c,
		Service: s,
	}
}
