package config

import (
	"os"
	"time"
)

const (
	EnvMasterDSN   = "ENGINE_DB_MASTER_DSN"
	EnvSlaveDSN    = "ENGINE_DB_SLAVE_DSN"
	EnvHTTPAddress = "ENGINE_HTTP_ADDR"

	HTTPAddrDefault = ":8081"
	DBDefaultDSN    = "postgres://postgres:postgres@localhost:5432/goclub?sslmode=disable"
)

type (
	DBDSN struct {
		Master string
		Slave  string
	}
	AppConfig struct {
		AppName     string
		AppVersion  string
		HTTPAddress string
		DSNs        DBDSN
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
	cfg.AppName = "GoClub"
	cfg.AppVersion = "v0.0.1"
	cfg.HTTPAddress = HTTPAddrDefault
	cfg.DSNs = DBDSN{Master: DBDefaultDSN, Slave: DBDefaultDSN}
}

func (cfg *AppConfig) updateWithEnv() {
	if env := os.Getenv(EnvHTTPAddress); env != "" {
		cfg.HTTPAddress = env
	}

	if env := os.Getenv(EnvMasterDSN); env != "" {
		cfg.DSNs.Master = env
	}

	if env := os.Getenv(EnvSlaveDSN); env != "" {
		cfg.DSNs.Slave = env
	}

}
