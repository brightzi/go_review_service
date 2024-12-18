// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"review_b/internal/biz"
	"review_b/internal/conf"
	"review_b/internal/data"
	"review_b/internal/server"
	"review_b/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, registry *conf.Registry, servicename *conf.Servicename, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	discovery := data.NewDiscovery(registry)
	reviewClient := data.NewReviewServiceClient(discovery, servicename)
	dataData, cleanup, err := data.NewData(confData, reviewClient, logger)
	if err != nil {
		return nil, nil, err
	}
	businessRepo := data.NewBusinessRepo(dataData, logger)
	businessUsecase := biz.NewBusinessUsecase(businessRepo, logger)
	businessService := service.NewBusinessService(businessUsecase)
	grpcServer := server.NewGRPCServer(confServer, businessService, logger)
	httpServer := server.NewHTTPServer(confServer, businessService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
