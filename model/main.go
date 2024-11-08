package main

import (
	"fmt"
	"goclub/model/common"
	"goclub/model/members"
	"reflect"
	"time"
)

func main() {
	bd := members.Birthday{Value: common.Date(time.Now().Truncate(time.Hour * 24)), Setted: true}
	m := members.Member{ID: 1, Name: "Alice", Birthday: bd, Notes: "qwerty"}
	foo(&m)
	fmt.Println(m)
}

func foo(x any) {
	rv := reflect.ValueOf(x)
	k := rv.Kind()

	switch k {
	case reflect.Pointer:
		e := rv.Elem()
		if e.CanInterface() {
			fmt.Println("interface -> ")
			foo(e.Interface())
		}
		for i := 0; i < e.NumField(); i++ {
			fld := e.Field(i)
			//fld := rv.Field(i).Interface()
			// if fld.IsValid() && fld.CanSet() {

			// }
			switch fld.Kind() {
			case reflect.String:
				fmt.Println("reflect.String")
				if fld.CanSet() {
					fmt.Println("CanSet!!!!")
					fld.SetString(fld.String() + "+")
				}
			case reflect.Int64:
				fmt.Println("reflect.Int64")
				if fld.CanSet() {
					fmt.Println("CanSet!!!!")
					fld.SetInt(fld.Int() + 1)
				}
			}

		}

	case reflect.Struct:
		t := reflect.TypeOf(x)
		for i := 0; i < t.NumField(); i++ {
			strctfld := t.Field(i)
			fld := rv.Field(i)
			var a any
			if strctfld.IsExported() {
				a = fld.Interface()
			}
			fmt.Println(" - field ", i, fld.Type(), fld.Kind(), " is ", strctfld, "|", strctfld.Name, "|", a)
			//fld := rv.Field(i).Interface()
			// if fld.IsValid() && fld.CanSet() {

			// }
			switch fld.Kind() {
			case reflect.String:
				fmt.Println("reflect.String")
				if fld.CanSet() {
					fmt.Println("CanSet!!!!")
					fld.SetString(fld.String() + "+")
				}
			case reflect.Int64:
				fmt.Println("reflect.Int64")
				if fld.CanSet() {
					fmt.Println("CanSet!!!!")
					fld.SetInt(fld.Int() + 1)
				}
			}

		}
	}

	fmt.Println(rv, k, "!!!!")
}
