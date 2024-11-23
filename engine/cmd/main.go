package main

import (
	"context"
	"fmt"
	"goclub/common/logger"
	"goclub/engine/internal/app"
	"goclub/engine/internal/config"
	"goclub/engine/internal/repository/db"
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

func main() {

	top := TopType{
		D: &PtrType{},
	}
	// x := db.DoExtractStructFields(reflect.TypeOf(top))
	// for k, v := range x {
	// 	fmt.Printf("%d: %+v\n", k, v)
	// }

	//ti :=	db.ExtractStructInfo(top)
	//fmt.Printf("%+v\n", ti)
	var q string
	var p1, p0 []any
	q = db.QueryTextForInsert(top)
	fmt.Println(q)
	p0 = []any{&top.A, &top.B, &top.C.E1, &top.C.E2, &top.C.E3.F1}
	if top.D != nil {
		p0 = append(p0, &top.D.P1, &top.D.P2)
	}
	p1 = db.QueryArgsForInsert(&top)
	fmt.Println(p0)
	fmt.Println(p1)
	// q = db.QueryTextForUpdate(top)
	// fmt.Println(q)
	// q = db.QueryTextForDelete(top)
	// fmt.Println(q)
	// q = db.QueryTextForRead(top)
	// fmt.Println(q)
	// q = db.QueryTextForList(top, nil)
	// fmt.Println(q)
	/*
		db.ExtractStructInfo(&top)
		db.DecodedTypes.Range(func(key, value any) bool {
			fmt.Println("Key:", key)
			fmt.Println("Val:", value)
			return true
		})
	*/

}

func main1() {
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
