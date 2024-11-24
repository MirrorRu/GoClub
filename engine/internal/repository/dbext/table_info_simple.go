package dbext

import (
	"goclub/model/common"
)

type (
	TableInfoSimple struct {
		name                                           string
		fields                                         map[int]common.TableField
		keyFieldNames, nonKeyFieldNames, allFieldNames []string
		keyFieldIds, nonKeyFieldIds, allFieldIds       []int
	}

	TableInfoProvider[T comparable] interface {
		Name() string
		Fields() map[int]common.TableField
		KeyFieldIds() []int
		NonKeyFieldIds() []int
		AllFieldIds() []int
		KeyFieldNames() []string
		NonKeyFieldNames() []string
		AllFieldNames() []string
		GetFieldsAddr(data *T, fieldIds []int) (pointers []any)
	}
)

func NewTableInfoSimple(name string,
	fields map[int]common.TableField,
	keyFieldIds, nonKeyFieldIds, allFieldIds []int) TableInfoSimple {
	return TableInfoSimple{
		name:             name,
		fields:           fields,
		keyFieldIds:      keyFieldIds,
		keyFieldNames:    common.CollectNames(fields, keyFieldIds),
		nonKeyFieldIds:   nonKeyFieldIds,
		nonKeyFieldNames: common.CollectNames(fields, nonKeyFieldIds),
		allFieldIds:      allFieldIds,
		allFieldNames:    common.CollectNames(fields, allFieldIds),
	}
}

func (ti *TableInfoSimple) Name() string {
	return ti.name
}
func (ti *TableInfoSimple) Fields() map[int]common.TableField {
	return ti.fields
}
func (ti *TableInfoSimple) KeyFieldIds() []int {
	return ti.keyFieldIds
}
func (ti *TableInfoSimple) NonKeyFieldIds() []int {
	return ti.nonKeyFieldIds
}
func (ti *TableInfoSimple) AllFieldIds() []int {
	return ti.allFieldIds
}
func (ti *TableInfoSimple) KeyFieldNames() []string {
	return ti.keyFieldNames
}
func (ti *TableInfoSimple) NonKeyFieldNames() []string {
	return ti.nonKeyFieldNames
}
func (ti *TableInfoSimple) AllFieldNames() []string {
	return ti.allFieldNames
}
