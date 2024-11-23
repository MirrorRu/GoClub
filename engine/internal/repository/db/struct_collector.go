package db

import (
	"reflect"
	"slices"
	"strings"
	"sync"
)

// pkgTag - ключ для получения аннотаций
const (
	maxStructDeep              = 16 // Максимальная глубина вложенности структур
	pkgTag                     = "db"
	pkgTagSeparator            = "|"
	pkgInTagSeparator          = ";"
	pkgTagGroupsKey            = "groups"
	pkgTagTitleKey             = "title"
	pkgTagNameKey              = "name"
	pkgTagTitleKeyValSeparator = "="
	pkgTagKeyGroupKey          = "key"   // группа полей первичного ключа
	pkgTagAutoGroupKey         = "auto"  // группа полей с генерацией внутри БД, обчыно автоматичеки генерируемые первичные ключи
	pkgTagEmbedKey             = "embed" // поля структуры с такой группой выделяются как отдельные единицы

)

type FieldInfo struct {
	Idx             int   // Основной индекс поля в списке
	StructFieldNums []int // Номер поля в структуре
	StructDeep      int   //Уровень вложенности структуры поля
	Name            string
	Title           string
	Type            reflect.Type
	Kind            reflect.Kind
	Groups          []string
	IsKey           bool
	IsAutogen       bool
}

type TableInfo struct {
	Name      string
	Type      reflect.Type
	Fields    []FieldInfo
	KeyFields []int
}

var DecodedTypes sync.Map  //Данные структур, по которым уже был сделан разбор
var DecodedFields sync.Map //Данные структур, по которым уже был сделан разбор

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
		Name:   srcType.Name(),
		Type:   srcType,
		Fields: DoExtractStructFields(srcType),
	}
	for idx, fld := range res.Fields {
		if fld.IsKey {
			res.KeyFields = append(res.KeyFields, idx)
		}

	}
	DecodedTypes.Store(srcType, res)
	return res
}

// func (ti *TableInfo) doExtractStructFields(t reflect.Type) {
// 	fmt.Println("doExtractStructFields", t)
// 	fldIdx := 0
// 	for i := 0; i < t.NumField(); i++ {
// 		strctfld := t.Field(i)
// 		if !strctfld.IsExported() {
// 			continue
// 		}

// 		fi := FieldInfo{
// 			Idx:  fldIdx,
// 			Name: strctfld.Name,
// 			Type: strctfld.Type,
// 			Kind: strctfld.Type.Kind(),
// 		}
// 		switch fi.Kind {
// 		case reflect.Struct:
// 			ptr := reflect.New(strctfld.Type).Interface()
// 			fmt.Println("CASE:", fi.Type, fi.Kind, reflect.TypeOf(ptr), reflect.ValueOf(ptr))
// 			ExtractStructInfo(ptr)
// 		case reflect.Pointer:
// 			elemType := strctfld.Type.Elem()
// 			ptr := reflect.New(elemType).Interface()
// 			fmt.Println("CASE:", elemType, fi.Type, fi.Kind, reflect.TypeOf(ptr), reflect.ValueOf(ptr))
// 			ExtractStructInfo(ptr)

// 		}

// 		tag := strctfld.Tag.Get(pkgTag)
// 		tagSplit := strings.Split(tag, pkgTagSeparator)
// 		for idx := range tagSplit {
// 			switch true {
// 			case strings.HasPrefix(tagSplit[idx], pkgTagGroupsKey+pkgTagTitleKeyValSeparator):
// 				s := strings.TrimPrefix(tagSplit[idx], pkgTagGroupsKey+pkgTagTitleKeyValSeparator)
// 				fi.Groups = strings.Split(s, pkgInTagSeparator)
// 				for _, g := range fi.Groups {
// 					switch g {
// 					case pkgTagKeyGroupKey:
// 						ti.KeyFields = append(ti.KeyFields, idx)
// 						fi.IsKey = true
// 					case pkgTagAutoGroupKey:
// 						fi.IsAutogen = true
// 					}
// 				}
// 			case strings.HasPrefix(tagSplit[idx], pkgTagTitleKey+pkgTagTitleKeyValSeparator):
// 				fi.Title = strings.TrimPrefix(tagSplit[idx], pkgTagTitleKey+pkgTagTitleKeyValSeparator)
// 			case strings.HasPrefix(tagSplit[idx], pkgTagNameKey+pkgTagTitleKeyValSeparator):
// 				fi.Name = strings.TrimPrefix(tagSplit[idx], pkgTagNameKey+pkgTagTitleKeyValSeparator)
// 			}
// 		}
// 		ti.Fields = append(ti.Fields, fi)
// 		//
// 		fldIdx++ // at last!

// 		// if fi.Kind == reflect.Struct {
// 		// 	ptr := reflect.New(fi.Type)
// 		// 	x := ptr.Elem().Interface()
// 		// 	doExtractStructInfo(x)
// 		// }
// 	}
// }

// doExtractStructFields - получение данных о полях таблицы соотвествующих данной структуре
// t - тип структуры
// startIdx - начальный индекс счетчика полей
func DoExtractStructFields(t reflect.Type) (result []FieldInfo) {
	saved, ok := DecodedFields.Load(t)
	if ok {
		return saved.([]FieldInfo)
	}
	fldIdx := 0
fldLoop:
	for numOfField := 0; numOfField < t.NumField(); numOfField++ {
		strctfld := t.Field(numOfField)
		if !strctfld.IsExported() {
			continue //Пропускаем неэкспортируемые поля
		}

		tag := strctfld.Tag.Get(pkgTag)
		tagSplit := strings.Split(tag, pkgTagSeparator) //Получили список тегов в аннотиции

		// Базовая инициализация данных поля
		fi := FieldInfo{
			Idx:             fldIdx,
			StructFieldNums: []int{numOfField},
			Name:            strctfld.Name,
			Title:           strctfld.Name,
			Type:            strctfld.Type,
			Kind:            strctfld.Type.Kind(),
		}

		// Уточнение свойств из аннотации
		for idx := range tagSplit {
			switch true {
			case tagSplit[idx] == pkgTagKeyGroupKey:
				fi.IsKey = true
			case strings.HasPrefix(tagSplit[idx], pkgTagGroupsKey+pkgTagTitleKeyValSeparator):
				s := strings.TrimPrefix(tagSplit[idx], pkgTagGroupsKey+pkgTagTitleKeyValSeparator)
				fi.Groups = strings.Split(s, pkgInTagSeparator)
				for _, g := range fi.Groups {
					switch g {
					case pkgTagKeyGroupKey:
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

		switch true {
		// Встраиваемая структура
		case fi.Kind == reflect.Struct && slices.Contains(tagSplit, pkgTagEmbedKey):
			tmp := DoExtractStructFields(strctfld.Type)
			titleSfx := ""
			if fi.Title != "" {
				titleSfx = " " + fi.Title
			}
			for _, v := range tmp {
				result = append(result, FieldInfo{
					Idx:             fldIdx,
					StructDeep:      v.StructDeep + 1,
					StructFieldNums: append([]int{numOfField}, v.StructFieldNums...),
					Name:            fi.Name + "_" + v.Name,
					Title:           v.Title + titleSfx,
					Type:            v.Type,
					Kind:            v.Kind,
				})
				fldIdx++
			}
			continue fldLoop
		// Указатель, т.е ключ другой таблицы
		case fi.Kind == reflect.Pointer:
			elemType := strctfld.Type.Elem()
			tmp := DoExtractStructFields(elemType)
			titleSfx := ""
			if fi.Title != "" {
				titleSfx = " " + fi.Title
			}
			for _, v := range tmp {
				if !v.IsKey {
					continue
				}
				result = append(result, FieldInfo{
					Idx:             fldIdx,
					StructDeep:      v.StructDeep + 1,
					StructFieldNums: append([]int{numOfField}, v.StructFieldNums...),
					Name:            fi.Name + "_" + v.Name,
					Title:           v.Title + titleSfx,
					Type:            v.Type,
					Kind:            v.Kind,
				})
				fldIdx++
			}
			continue fldLoop
		}

		result = append(result, fi)
		//
		fldIdx++ // at last!
	}
	DecodedTypes.Store(t, result)
	return result
}
