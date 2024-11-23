package dbext

import (
	"goclub/model/common"
	"strconv"
	"strings"
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

func QueryTextForInsert(tableName string, insertFieldsIds, returnFieldsIds []string) string {
	fields := make([]string, len(insertFieldsIds))
	for idx := range insertFieldsIds {
		fields[idx] = "$" + strconv.Itoa(idx+1)
	}
	q := "insert into " +
		tableName +
		" (" +
		strings.Join(insertFieldsIds, ", ") +
		") values (" +
		strings.Join(fields, ", ") +
		") returning " +
		strings.Join(returnFieldsIds, ", ")
	return q
}

func QueryTextForDelete(tableName string, keyFieldIds []string) string {
	keys := make([]string, len(keyFieldIds))
	for idx, id := range keyFieldIds {
		keys[idx] = id + "=$" + strconv.Itoa(idx+1)
	}
	q := "delete from " +
		tableName +
		" where " +
		strings.Join(keys, " and ")
	return q
}

func QueryTextForUpdate(tableName string, updateFieldIds, keyFieldIds, returnFieldIds []string) string {
	fields := make([]string, len(updateFieldIds))
	for idx, id := range updateFieldIds {
		fields[idx] = id + "=$" + strconv.Itoa(idx+1)
	}
	keys := make([]string, len(keyFieldIds))
	for idx, id := range keyFieldIds {
		keys[idx] = id + "=$" + strconv.Itoa(len(fields)+idx+1)
	}
	q := "update " +
		tableName +
		" set " +
		strings.Join(fields, ", ") +
		" where " +
		strings.Join(keys, " and ") +
		" returning " +
		strings.Join(returnFieldIds, ", ")
	return q
}

func QueryTextForRead(tableName string, keyFieldIds, returnFieldIds []string) string {
	keys := make([]string, len(keyFieldIds))
	for idx, id := range keyFieldIds {
		keys[idx] = id + "=$" + strconv.Itoa(idx+1)
	}
	q := "select " +
		strings.Join(returnFieldIds, ", ") +
		" from " +
		tableName +
		" where " +
		strings.Join(keys, " and ")
	return q
}

func QueryTextForList(tableName string, keyFieldIds, returnFieldIds []string, filter common.ListFilter) string {
	q := "select " +
		strings.Join(returnFieldIds, ", ") +
		" from " +
		tableName +
		" order by " +
		strings.Join(keyFieldIds, ", ")
	return q
}
