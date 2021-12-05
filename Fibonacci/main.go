package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/rauf95/rauf/api/rest"
	"github.com/rauf95/rauf/api/rpc"
)

var (
	cfg = flag.String("config", "config.json", "sets config")
)

type Config struct {
	HTTPPort int `json:"http_port"`
	GRPCPort int `json:"grpc_port"`
}

func main() {
	flag.Parse()
	logger := zerolog.New(os.Stdout).With().Caller().Timestamp().Logger()
	ctxParent := logger.WithContext(context.Background())

	err := start(ctxParent)
	if err != nil {
		logger.Fatal().Err(err).Msg("fatal error")
	}
}

func start(ctx context.Context) error {
	logger := zerolog.Ctx(ctx)

	if cfg == nil || *cfg == "" {
		return errors.New("config must be set")
	}

	cfgFile, err := os.Open(*cfg)
	if err != nil {
		return fmt.Errorf("os.Open: %w", err)
	}

	serverConfig := Config{}
	err = json.NewDecoder(cfgFile).Decode(&serverConfig)
	if err != nil {
		return fmt.Errorf("json.NewDecoder.Decode: %w", err)
	}

	grpcAPI := rpc.New()
	httpAPI := rest.New(*logger)

	g, ctx := errgroup.WithContext(ctx)
	g.Go(runGRPC(ctx, "0.0.0.0", serverConfig.GRPCPort, grpcAPI))
	g.Go(runHTTP(ctx, "0.0.0.0", serverConfig.HTTPPort, httpAPI))

	return g.Wait()
}

func runGRPC(ctx context.Context, host string, port int, srv *grpc.Server) func() error {
	return func() error {
		logger := zerolog.Ctx(ctx).With().Str("name", "grpc").Logger()
		ln, err := net.Listen("tcp", net.JoinHostPort(host, strconv.Itoa(port)))
		if err != nil {
			return fmt.Errorf("net.Listen: %w", err)
		}

		logger.Info().Str("host", host).Int("port", port).Msg("started")
		defer logger.Info().Msg("shutdown")

		return srv.Serve(ln)
	}
}

func runHTTP(ctx context.Context, host string, port int, handler http.Handler) func() error {
	return func() error {
		logger := zerolog.Ctx(ctx).With().Str("name", "http").Logger()

		srv := &http.Server{
			Addr:    net.JoinHostPort(host, strconv.Itoa(port)),
			Handler: handler,
		}

		logger.Info().Str("host", host).Int("port", port).Msg("started")
		defer logger.Info().Msg("shutdown")

		return srv.ListenAndServe()
	}
}
