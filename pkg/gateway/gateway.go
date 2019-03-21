package gateway

import (
	"context"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	gw "github.com/amber-lamp/go-grpc-sample/pkg/api"
)

type Config struct {
	Server string
	Port   string
}

func RunServer() error {

	cfg := getConfig()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterBookServiceHandlerFromEndpoint(ctx, mux, cfg.Server, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":"+cfg.Port, mux)
}

func getConfig() Config {
	var cfg Config

	cfg.Server = os.Getenv("GRPC_SERVER")
	cfg.Port = os.Getenv("GATEWAY_PORT")

	return cfg
}
