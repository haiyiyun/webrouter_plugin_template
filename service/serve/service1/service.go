package service1

import (
	"github.com/haiyiyun/webrouter_plugin_template/service/serve"
)

type Service struct {
	*serve.Service
}

func NewService(s *serve.Service) *Service {
	return &Service{
		Service: s,
	}
}
