package config

import (
	"os"
	"time"
)

const (
	EnvMasterDSN   = "GEARS_DB_MASTER_DSN"
	EnvSlaveDSN    = "GEARS_DB_SLAVE_DSN"
	EnvHTTPAddress = "GEARS_HTTP_ADDR"
	EnvGRPCAddress = "GEARS_GRPC_ADDR"

	HTTPAddrDefault = ":8081"
	GRPCAddrDefault = ":8082"
	DBDefaultDSN    = "postgres://postgres:postgres@localhost:5432/goclub?sslmode=disable"
)

type (
	DSNPair struct {
		Master string
		Slave  string
	}
	AppConfig struct {
		AppMetricsID string
		AppTraceID   string
		JaegerURL    string

		HTTPAddress string
		GRPCAddress string
		DSNs        DSNPair
	}
)

var (
	DBTimeout time.Duration = 2 * time.Second
	Cfg       *AppConfig    = config()
)

func config() *AppConfig {
	cfg := &AppConfig{}
	cfg.initDefault()
	cfg.updateWithEnv()
	return cfg
}

func (cfg *AppConfig) initDefault() {
	cfg.GRPCAddress = GRPCAddrDefault
	cfg.HTTPAddress = HTTPAddrDefault
	cfg.DSNs = DSNPair{Master: DBDefaultDSN, Slave: DBDefaultDSN}
}

func (cfg *AppConfig) updateWithEnv() {
	if env := os.Getenv(EnvHTTPAddress); env != "" {
		cfg.HTTPAddress = env
	}
	if env := os.Getenv(EnvGRPCAddress); env != "" {
		cfg.GRPCAddress = env
	}

	if env := os.Getenv(EnvMasterDSN); env != "" {
		cfg.DSNs.Master = env
	}

	if env := os.Getenv(EnvSlaveDSN); env != "" {
		cfg.DSNs.Slave = env
	}

}
