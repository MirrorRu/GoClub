package main

import (
	"context"
	"fmt"
	"goclub/common/logger"
	"goclub/engine/internal/app"
	"goclub/engine/internal/config"
	"goclub/engine/internal/repository/db"
	"goclub/model/common"
	"goclub/model/members"
	"os"
	"time"
)

func main() {
	bd := members.Birthday{Value: common.Date(time.Now().Truncate(time.Hour * 24)), Setted: true}
	m := members.Member{ID: 1, Name: "Alice", Birthday: bd, Notes: "qwerty"}

	/*
		fmt.Println(reflect.TypeOf(m))
		fmt.Println(reflect.TypeOf(&m))
	*/
	efs := db.ExtractStructInfo(&m)
	fmt.Println(efs)
	efs = db.ExtractStructInfo(m)
	fmt.Println(efs)

	db.DecodedTypes.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})

	// foo(&m)
	// fmt.Println(m)
}
func main0() {
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
