package db

import (
	"reflect"
	"strconv"
	"strings"
	"sync"
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

var SavedQuerys sync.Map

type savedQueryKey struct {
	ObjType   reflect.Type
	QueryType QueryType
}

func QueryTextForRead(record any) (query string) {
	tableInfo := ExtractStructInfo(record)
	saveKey := savedQueryKey{ObjType: tableInfo.Type, QueryType: QueryTypeSelect}
	val, ok := SavedQuerys.Load(saveKey)
	if ok {
		return val.(string)
	}

	sb := strings.Builder{}
	sb.WriteString("select ")
	for idx := range tableInfo.Fields {
		if idx > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(tableInfo.Fields[idx].Name)
	}
	sb.WriteString(" from " + tableInfo.Name + " where ")

	for idx := range tableInfo.KeyFields {
		fIdx := tableInfo.KeyFields[idx]
		if idx > 0 {
			sb.WriteString(" AND ")
		}
		sb.WriteString(tableInfo.Fields[fIdx].Name + "=$" + strconv.Itoa(idx+1))
	}
	query = sb.String()
	SavedQuerys.Store(saveKey, query)

	return query
}

func QueryArgsForRead(record any) (pointers []any) {
	tableInfo := ExtractStructInfo(record)

	/*
		t := reflect.TypeOf(model).Elem()
		v := reflect.ValueOf(model).Elem()
		m := make(map[string]interface{})
		for i := 0; i < v.NumField(); i {
			colName := t.Field(i).Tag.Get("sql")
			field := v.Field(i)
			m[colName] = field.Addr().Interface()
		}
		return m
	*/
	v := reflect.ValueOf(record).Elem()
	pointers = make([]any, len(tableInfo.KeyFields))
	for idx := range tableInfo.KeyFields {
		fIdx := tableInfo.KeyFields[idx]
		pointers[idx] = v.Field(fIdx).Addr().Interface()
	}
	return pointers
}

func ScanArgsForRead(record any) (pointers []any) {
	tableInfo := ExtractStructInfo(record)

	v := reflect.ValueOf(record).Elem()
	pointers = make([]any, len(tableInfo.Fields))
	for idx := range tableInfo.Fields {
		pointers[idx] = v.Field(idx).Addr().Interface()
	}
	return pointers
}

func QueryTextForInsert(record any) (query string) {
	tableInfo := ExtractStructInfo(record)
	saveKey := savedQueryKey{ObjType: tableInfo.Type, QueryType: QueryTypeInsert}
	val, ok := SavedQuerys.Load(saveKey)
	if ok {
		return val.(string)
	}

	fields4insert := make([]int, 0, len(tableInfo.Fields))
	fields4return := make([]int, 0, len(tableInfo.Fields))

	sb := strings.Builder{}
	sb.WriteString("insert into " + tableInfo.Name + "(")
	for idx := range tableInfo.Fields {
		if tableInfo.Fields[idx].IsAutogen {
			fields4return = append(fields4return, idx)
			continue
		}

		if len(fields4insert) > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(tableInfo.Fields[idx].Name)
		fields4insert = append(fields4insert, idx)
	}

	sb.WriteString(") values (")
	for idx := range fields4insert {
		if idx > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString("$" + strconv.Itoa(idx+1))
	}

	sb.WriteString(") returning ")
	for idx := range tableInfo.Fields {
		if idx > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(tableInfo.Fields[idx].Name)
	}

	query = sb.String()
	SavedQuerys.Store(saveKey, query)

	return query
}

func QueryArgsForInsert(record any) (pointers []any) {
	tableInfo := ExtractStructInfo(record)

	v := reflect.ValueOf(record).Elem()
	pointers = make([]any, 0, len(tableInfo.KeyFields))
	for idx := range tableInfo.Fields {
		if tableInfo.Fields[idx].IsAutogen {
			continue
		}
		pointers = append(pointers, v.Field(idx).Addr().Interface())
	}
	return pointers
}

func ScanArgsForAllFields(record any) (pointers []any) {
	tableInfo := ExtractStructInfo(record)

	v := reflect.ValueOf(record).Elem()
	pointers = make([]any, len(tableInfo.Fields))
	for idx := range tableInfo.Fields {
		pointers[idx] = v.Field(idx).Addr().Interface()
	}
	return pointers
}

func QueryTextForUpdate(record any) (query string) {
	tableInfo := ExtractStructInfo(record)
	saveKey := savedQueryKey{ObjType: tableInfo.Type, QueryType: QueryTypeUpdate}
	val, ok := SavedQuerys.Load(saveKey)
	if ok {
		return val.(string)
	}

	fields4update := make([]int, 0, len(tableInfo.Fields))

	sb := strings.Builder{}
	sb.WriteString("update " + tableInfo.Name + " set ")
	cnt := 0
	for idx := range tableInfo.Fields {
		if tableInfo.Fields[idx].IsKey {
			continue
		}
		cnt++
		if cnt > 1 {
			sb.WriteString(", ")
		}
		sb.WriteString(tableInfo.Fields[idx].Name + "=$" + strconv.Itoa(cnt))
		fields4update = append(fields4update, idx)
	}

	sb.WriteString(" where ")
	for idx := range tableInfo.KeyFields {
		if idx > 0 {
			sb.WriteString(" AND ")
		}
		cnt++
		sb.WriteString(tableInfo.Fields[tableInfo.KeyFields[idx]].Name + "=$" + strconv.Itoa(cnt))
	}

	sb.WriteString(" returning ")
	for idx := range tableInfo.Fields {
		if idx > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(tableInfo.Fields[idx].Name)
	}

	query = sb.String()
	SavedQuerys.Store(saveKey, query)

	return query
}

func QueryArgsForUpdate(record any) (pointers []any) {
	tableInfo := ExtractStructInfo(record)

	v := reflect.ValueOf(record).Elem()
	pointers = make([]any, 0, len(tableInfo.Fields))
	for idx := range tableInfo.Fields {
		if tableInfo.Fields[idx].IsKey {
			continue
		}
		pointers = append(pointers, v.Field(idx).Addr().Interface())
	}

	for idx := range tableInfo.KeyFields {
		pointers = append(pointers, v.Field(tableInfo.KeyFields[idx]).Addr().Interface())
	}

	return pointers
}

func QueryTextForList(record any, filter any) (query string) {
	tableInfo := ExtractStructInfo(record)
	saveKey := savedQueryKey{ObjType: tableInfo.Type, QueryType: QueryTypeListing}
	val, ok := SavedQuerys.Load(saveKey)
	if ok {
		return val.(string)
	}

	sb := strings.Builder{}
	sb.WriteString("select ")
	for idx := range tableInfo.Fields {
		if idx > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(tableInfo.Fields[idx].Name)
	}
	sb.WriteString(" from " + tableInfo.Name)

	query = sb.String()
	SavedQuerys.Store(saveKey, query)

	return query
}

func QueryArgsForList(record any, filter any) (pointers []any) {
	return pointers
}

func QueryTextForDelete(record any) (query string) {
	tableInfo := ExtractStructInfo(record)
	saveKey := savedQueryKey{ObjType: tableInfo.Type, QueryType: QueryTypeDelete}
	val, ok := SavedQuerys.Load(saveKey)
	if ok {
		return val.(string)
	}

	sb := strings.Builder{}
	sb.WriteString("delete from " + tableInfo.Name + " where ")
	cnt := 0
	for idx := range tableInfo.KeyFields {
		if idx > 0 {
			sb.WriteString(" AND ")
		}
		cnt++
		sb.WriteString(tableInfo.Fields[tableInfo.KeyFields[idx]].Name + "=$" + strconv.Itoa(cnt))
	}

	query = sb.String()
	SavedQuerys.Store(saveKey, query)

	return query
}

func QueryArgsForDelete(record any) (pointers []any) {
	tableInfo := ExtractStructInfo(record)

	v := reflect.ValueOf(record).Elem()
	pointers = make([]any, 0, len(tableInfo.KeyFields))

	for idx := range tableInfo.KeyFields {
		pointers = append(pointers, v.Field(tableInfo.KeyFields[idx]).Addr().Interface())
	}

	return pointers
}
