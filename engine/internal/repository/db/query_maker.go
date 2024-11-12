package db

import (
	"reflect"
	"strings"
	"sync"
)

// pkgTag - ключ для получения аннотаций
const (
	pkgTag             = "db"
	pkgTagSeparator    = ";"
	pkgInTagSeparator  = ","
	pkgTagGroupsPrefix = "groups="
)

// ТИп запроса - Вставка, обновление, удаление, единичное чтение и пр.
type QueryType int

const (
	QueryTypeSelect QueryType = iota
	QueryTypeListing
	QueryTypeInsert
	QueryTypeUpdate
	QueryTypeDelete
)

type FieldInfo struct {
	Idx    int
	Name   string
	Type   reflect.Type
	Kind   reflect.Kind
	Groups []string
}

type TableInfo struct {
	Name   string
	Fields []FieldInfo
}

var DecodedTypes sync.Map //Данные структур, по которым уже был сделан разбор

func ExtractStructInfo(record any) *TableInfo {
	rv := reflect.ValueOf(record)
	k := rv.Kind()
	switch k {
	case reflect.Pointer:
		e := rv.Elem()
		if e.CanInterface() {
			return doExtractStructInfo(e.Interface())
		}
	case reflect.Struct:
		return doExtractStructInfo(record)
	}
	return nil

}

func doExtractStructInfo(src any) (res *TableInfo) {
	t := reflect.TypeOf(src)
	saved, ok := DecodedTypes.Load(t)
	if ok {
		return saved.(*TableInfo)
	}
	rv := reflect.ValueOf(src)
	k := rv.Kind()
	if k != reflect.Struct {
		return res
	}

	res = &TableInfo{
		Name:   t.Name(),
		Fields: doExtractStructFields(t),
	}
	DecodedTypes.Store(t, res)
	return res
}

func doExtractStructFields(t reflect.Type) (res []FieldInfo) {
	res = make([]FieldInfo, 0, t.NumField())
	fldIdx := 0
	for i := 0; i < t.NumField(); i++ {
		strctfld := t.Field(i)
		if !strctfld.IsExported() {
			continue
		}
		fi := FieldInfo{
			Idx:  fldIdx,
			Name: strctfld.Name,
			Type: strctfld.Type,
			Kind: strctfld.Type.Kind(),
		}

		tag := strctfld.Tag.Get(pkgTag)
		tagSplit := strings.Split(tag, pkgTagSeparator)
		for i := range tagSplit {
			switch true {
			case strings.HasPrefix(tagSplit[i], pkgTagGroupsPrefix):
				s := strings.TrimPrefix(tagSplit[i], pkgTagGroupsPrefix)
				fi.Groups = strings.Split(s, pkgInTagSeparator)
			}
		}
		res = append(res, fi)
		//
		fldIdx++ // at last!

		if fi.Kind == reflect.Struct {
			ptr := reflect.New(fi.Type)
			x := ptr.Elem().Interface()
			doExtractStructInfo(x)
		}
	}
	return res
}

// func QueryForSelect(record any) (query string, err error) {

// }
