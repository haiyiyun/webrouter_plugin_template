package service1

import (
	"github.com/haiyiyun/webrouter_plugin_template/service"
)

type Service struct {
	*service.Service
}

func NewService(s *service.Service) *Service {
	return &Service{
		Service: s,
	}
}
