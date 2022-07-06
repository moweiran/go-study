package svc

import (
	"app/repository"
	"app/svc/iface"
)

type ServiceContext struct {
	Provider *repository.Provider
	Service  iface.IService
}

var ServiceComm *ServiceContext

func NewServiceContext() *ServiceContext {
	p := repository.NewProvider()
	ServiceComm = &ServiceContext{
		Provider: p,
	}
	return ServiceComm
}
