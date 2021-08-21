package service1

import (
	"github.com/haiyiyun/webrouter_plugin_template/service/manage"
)

type Service struct {
	*manage.Service
}

func NewService(s *manage.Service) *Service {
	return &Service{
		Service: s,
	}
}
