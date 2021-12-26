// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/mohuishou/go-training/Week04/homework/internal/repository"
	"github.com/mohuishou/go-training/Week04/homework/internal/service"
	"github.com/mohuishou/go-training/Week04/homework/internal/usecase"
)

// Injectors from wire.go:

func serviceInit() (*service.UserService, error) {
	viper, err := InitConfig()
	if err != nil {
		return nil, err
	}
	client, err := NewDB(viper)
	if err != nil {
		return nil, err
	}
	iUserRepo := repository.NewRepository(client)
	iUserUsecase := usecase.NewUserUsecase(iUserRepo)
	userService := service.NewUserService(iUserUsecase)
	return userService, nil
}
