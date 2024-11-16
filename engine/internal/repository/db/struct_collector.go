package db

import (
	"reflect"
	"strings"
	"sync"
)

// pkgTag - ключ для получения аннотаций
const (
	pkgTag                     = "db"
	pkgTagSeparator            = "|"
	pkgInTagSeparator          = ";"
	pkgTagGroupsKey            = "groups"
	pkgTagTitleKey             = "title"
	pkgTagNameKey              = "name"
	pkgTagTitleKeyValSeparator = "="
	pkgTagKeyGroupKey          = "key"  // группа полей первичного ключа
	pkgTagAutoGroupKey         = "auto" // группа полей с генерацией внутри БД, обчыно автоматичеки генерируемые первичные ключи
)

type FieldInfo struct {
	Idx       int
	Name      string
	Title     string
	Type      reflect.Type
	Kind      reflect.Kind
	Groups    []string
	IsKey     bool
	IsAutogen bool
}

type TableInfo struct {
	Name      string
	Type      reflect.Type
	Fields    []FieldInfo
	KeyFields []int
}

var DecodedTypes sync.Map //Данные структур, по которым уже был сделан разбор

func ExtractStructInfo(record any) *TableInfo {
	rv := reflect.ValueOf(record)
	k := rv.Kind()
	switch k {
	case reflect.Pointer:
		e := rv.Elem()
		if e.CanInterface() {
			return newTableInfo(e.Interface())
		}
	case reflect.Struct:
		return newTableInfo(record)
	}
	return nil

}

func newTableInfo(src any) (res *TableInfo) {
	srcType := reflect.TypeOf(src)
	saved, ok := DecodedTypes.Load(srcType)
	if ok {
		return saved.(*TableInfo)
	}
	rv := reflect.ValueOf(src)
	k := rv.Kind()
	if k != reflect.Struct {
		return res
	}

	res = &TableInfo{
		Name: srcType.Name(),
		Type: srcType,
	}
	res.doExtractStructFields(srcType)
	DecodedTypes.Store(srcType, res)
	return res
}

func (ti *TableInfo) doExtractStructFields(t reflect.Type) {
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
		for idx := range tagSplit {
			switch true {
			case strings.HasPrefix(tagSplit[idx], pkgTagGroupsKey+pkgTagTitleKeyValSeparator):
				s := strings.TrimPrefix(tagSplit[idx], pkgTagGroupsKey+pkgTagTitleKeyValSeparator)
				fi.Groups = strings.Split(s, pkgInTagSeparator)
				for _, g := range fi.Groups {
					switch g {
					case pkgTagKeyGroupKey:
						ti.KeyFields = append(ti.KeyFields, idx)
						fi.IsKey = true
					case pkgTagAutoGroupKey:
						fi.IsAutogen = true
					}
				}
			case strings.HasPrefix(tagSplit[idx], pkgTagTitleKey+pkgTagTitleKeyValSeparator):
				fi.Title = strings.TrimPrefix(tagSplit[idx], pkgTagTitleKey+pkgTagTitleKeyValSeparator)
			case strings.HasPrefix(tagSplit[idx], pkgTagNameKey+pkgTagTitleKeyValSeparator):
				fi.Name = strings.TrimPrefix(tagSplit[idx], pkgTagNameKey+pkgTagTitleKeyValSeparator)
			}
		}
		ti.Fields = append(ti.Fields, fi)
		//
		fldIdx++ // at last!

		// if fi.Kind == reflect.Struct {
		// 	ptr := reflect.New(fi.Type)
		// 	x := ptr.Elem().Interface()
		// 	doExtractStructInfo(x)
		// }
	}
}
