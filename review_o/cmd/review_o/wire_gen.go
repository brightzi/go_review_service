// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"review_o/internal/biz"
	"review_o/internal/conf"
	"review_o/internal/data"
	"review_o/internal/server"
	"review_o/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, registry *conf.Registry, servicename *conf.Servicename, logger log.Logger) (*kratos.App, func(), error) {
	discovery := data.NewDiscovery(registry)
	reviewClient := data.NewReviewServiceClient(discovery, servicename)
	dataData, cleanup, err := data.NewData(confData, reviewClient, logger)
	if err != nil {
		return nil, nil, err
	}
	operationRepo := data.NewOperationRepo(dataData, logger)
	operationUsecase := biz.NewOperationUsecase(operationRepo, logger)
	operationService := service.NewOperationService(operationUsecase)
	grpcServer := server.NewGRPCServer(confServer, operationService, logger)
	httpServer := server.NewHTTPServer(confServer, operationService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
