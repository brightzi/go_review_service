package data

import (
	"context"
	"fmt"
	v1 "review_b/api/review/v1"
	"review_b/internal/conf"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"github.com/hashicorp/consul/api"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewDiscovery, NewData, NewBusinessRepo, NewReviewServiceClient)

// Data .
type Data struct {
	// TODO wrapped database client
	rc  v1.ReviewClient
	log *log.Helper
}

// NewData .
func NewData(c *conf.Data, rc v1.ReviewClient, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{rc: rc, log: log.NewHelper(logger)}, cleanup, nil
}

func NewReviewServiceClient(d registry.Discovery, conf *conf.Servicename) v1.ReviewClient {
	address := fmt.Sprintf("discovery:///%s", conf.Servername)
	conn, err := grpc.DialInsecure(
		context.Background(),
		// grpc.WithEndpoint("127.0.0.1:9000"),
		grpc.WithEndpoint(address),
		grpc.WithDiscovery(d),
		grpc.WithMiddleware(
			recovery.Recovery(),
			validate.Validator(),
		),
	)
	if err != nil {
		panic(err)
	}
	return v1.NewReviewClient(conn)
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := api.DefaultConfig()
	c.Address = conf.Consul.Address // 使用配置文件中的值
	c.Scheme = conf.Consul.Scheme

	client, err := api.NewClient(c)
	if err != nil {
		panic(err)
	}
	// new reg with consul client
	dis := consul.New(client)
	return dis
}
