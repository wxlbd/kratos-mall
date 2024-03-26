//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/wxlbd/kratos-pms/internal/biz"
	"github.com/wxlbd/kratos-pms/internal/conf"
	"github.com/wxlbd/kratos-pms/internal/data"
	"github.com/wxlbd/kratos-pms/internal/server"
	"github.com/wxlbd/kratos-pms/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	tintlog "github.com/wxlbd/kit/log"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *tintlog.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		wire.Bind(new(log.Logger), new(*tintlog.Logger)),
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		newApp,
	))
}
