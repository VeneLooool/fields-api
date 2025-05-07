package config

import (
	"context"
	"os"
)

const (
	EnvKeyHttpPort = "FIELDS_HTTP_PORT"
	EnvKeyGrpcPort = "FIELDS_GRPC_PORT"
)

type Config struct {
	HttpPort string
	GrpcPort string
}

func New(ctx context.Context) (*Config, error) {
	return &Config{
		HttpPort: os.Getenv(EnvKeyHttpPort),
		GrpcPort: os.Getenv(EnvKeyGrpcPort),
	}, nil
}
