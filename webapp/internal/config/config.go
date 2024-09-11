package config

import (
	"context"
	"os"
)

const (
	appTitleDflt      = "Go Club"
	appTitleEnv       = "GOCLUB_APP_TITLE"
	databaseDSNDflt   = "postgresql://postgress:postgress@localhost:5432/goclub?sslmode=disable"
	databaseDSNEnv    = "GOCLUB_DB_DSN"
	serverAddressDflt = ":8088"
	serverAddressEnv  = "GOCLUB_SRV_ADDRESS"
)

type (
	AppConfig struct {
		Ctx           context.Context
		AppTitle      string
		DatabaseDSN   string
		ServerAddress string
	}
)

var Cfg = loadAppConfig()

func loadAppConfig() *AppConfig {
	cfg := &AppConfig{
		Ctx:           context.Background(),
		AppTitle:      appTitleDflt,
		DatabaseDSN:   databaseDSNDflt,
		ServerAddress: serverAddressDflt,
	}

	if env := os.Getenv(appTitleEnv); env != "" {
		cfg.AppTitle = env
	}

	if env := os.Getenv(databaseDSNEnv); env != "" {
		cfg.DatabaseDSN = env
	}

	if env := os.Getenv(serverAddressEnv); env != "" {
		cfg.ServerAddress = env
	}

	return cfg
}
