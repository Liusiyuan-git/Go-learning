//+build wireinject

package main

import (
	"github.com/google/wire"
)

func serviceInit() (*service.UserService, error) {
	wire.Build(UserSet, NewDB, InitConfig)
	return new(service.UserService), nil
}
