package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	tintlog "github.com/wxlbd/kit/log"
	"github.com/wxlbd/kratos-pms/internal/conf"
	clientv3 "go.etcd.io/etcd/client/v3"
	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(c *conf.Data, logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	// new etcd client
	client, err := clientv3.New(clientv3.Config{
		Endpoints: c.Etcd.Addresses,
		Username:  c.Etcd.Username,
		Password:  c.Etcd.Password,
	})
	if err != nil {
		panic(err)
	}
	// new reg with etcd client
	reg := etcd.New(client)
	return kratos.New(
		kratos.ID(id),
		kratos.Name("pms"),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
		kratos.Registrar(reg),
	)
}

func main() {
	flag.Parse()
	//logger := log.With(log.NewStdLogger(os.Stdout),
	//	"ts", log.DefaultTimestamp,
	//	"caller", log.DefaultCaller,
	//	"service.id", id,
	//	"service.name", Name,
	//	"service.version", Version,
	//	"trace.id", tracing.TraceID(),
	//	"span.id", tracing.SpanID(),
	//)

	logger := tintlog.NewLogger(tintlog.WithWriter(os.Stdout), tintlog.WithLevel(slog.LevelDebug))
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
