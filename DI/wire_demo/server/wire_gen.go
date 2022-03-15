// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/micro/go-micro/v2"
	"wireDemo/server/handle"
	"wireDemo/server/model"
)

// Injectors from wire.go:

func initApp() (micro.Service, error) {
	modelModel, err := model.NewModel()
	if err != nil {
		return nil, err
	}
	handleHandle := handle.NewHandle(modelModel)
	service, err := handle.NewService(handleHandle)
	if err != nil {
		return nil, err
	}
	return service, nil
}