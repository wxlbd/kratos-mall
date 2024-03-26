package server

import (
	"context"

	v1 "kratos-admin/api/product/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"kratos-admin/internal/biz"
	"kratos-admin/internal/conf"
	"kratos-admin/internal/service"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, productSvc *service.ProductService, logger log.Logger, au *biz.AuthUseCase) *http.Server {
	opts := []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),    // 异常恢复
			tracing.Server(),       // 链路追踪中间件
			logging.Server(logger), // 日志中间件
			metrics.Server(),       // 监控中间件
			validate.Validator(),   // 参数校验
			metadata.Server(),      // 元数据中间
			// selector.Server(au.JwtMiddleware()).
			//	Match(NewWhiteListMatcher()).Build(), // 路由选择中间件
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterProductServiceHTTPServer(srv, productSvc)
	return srv
}

// NewWhiteListMatcher 路由白名单
func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList["/shop.interface.v1.ShopInterface/Login"] = struct{}{}
	whiteList["/shop.interface.v1.ShopInterface/Register"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}
