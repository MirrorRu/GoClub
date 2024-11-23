package main

import (
	"context"
	"fmt"
	"goclub/common/logger"
	"goclub/engine/internal/app"
	"goclub/engine/internal/config"
	"os"
)

type InclType struct {
	I1, I2 int
}

type EmbType struct {
	E1, E2 int
	E3     DeepEmbType `db:"embed"`
}

type DeepEmbType struct {
	F1 int
}

type PtrType struct {
	P1 int `db:"groups=key"`
	P2 int `db:"key"`
	P3 int
}

type TopType struct {
	A int `db:"groups=key"`
	B InclType
	C EmbType `db:"embed|title=Сложное полюшко"`
	D *PtrType
	E *PtrType
}

func main0() {
	// tr := tarifrepo.NewTarifsRepo(nil, nil)
	// fmt.Println(tr.QueryTextForInsert())
	// fmt.Println(tr.QueryTextForUpdate())
	// fmt.Println(tr.QueryTextForDelete())
	// fmt.Println(tr.QueryTextForRead())
	// fmt.Println(tr.QueryTextForList(nil))
}

func main() {
	cfg := config.Cfg
	ctx := context.Background()

	logger.Info(ctx, fmt.Sprintf("%s %s is starting", cfg.AppName, cfg.AppVersion))

	app := app.NewApp()
	if err := app.Run(ctx, cfg); err != nil {
		logger.Error(ctx, "Application finished with error", err)
		os.Exit(-1)
	}
	logger.Info(ctx, "Application is finished")
}
