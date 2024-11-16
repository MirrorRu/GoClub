package main

import (
	"context"
	"fmt"
	"goclub/common/logger"
	"goclub/engine/internal/app"
	"goclub/engine/internal/config"
	"goclub/engine/internal/repository/db"
	"goclub/model/members"
	"os"
	"time"
)

func main0() {
	//bd := members.Birthday{Value: common.Date(time.Now().Truncate(time.Hour * 24)), Setted: true}
	now := time.Now()
	bd := now.Year()*10000 + int(now.Month())*100 + now.Day()
	m := members.Member{ID: 1, FullName: "Alice", Birthday: members.Birthday(bd), Notes: "qwerty"}

	/*
		fmt.Println(reflect.TypeOf(m))
		fmt.Println(reflect.TypeOf(&m))
	*/
	efs := db.ExtractStructInfo(&m)
	fmt.Println(efs)
	var s string
	s = db.QueryTextForRead(m)
	fmt.Println(s)
	s = db.QueryTextForRead(&m)
	fmt.Println(s)
	ptrs := db.QueryArgsForRead(&m)
	fmt.Println(ptrs)
	ptrs = db.QueryArgsForRead(m)
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
