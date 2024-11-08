package common

import "time"

type (
	ID int64

	Name  string
	Notes string
	Date  time.Time

	Option[T any] struct {
		Setted bool // Is value setted and valid
		Value  T
	}

	Result[T any] struct {
		Value T
		Err   error
	}

	FieldType uint

	DataFieldDefinition struct {
		Type  FieldType
		Name  string
		Title string
	}

	DataFieldValue any

	DataTableDefinion struct {
		Name   string
		Title  string
		Fields []DataFieldDefinition
	}

	DataTable interface {
		//Перечисление всех полей таблицы
		Definition() *DataTableDefinion

		// Ключевые (идентифицирующие) поля
		KeyFields() []DataFieldDefinition

		// Основные поля, несущие самое важное
		BasicFields() []DataFieldDefinition

		// Дополнительные поля со вспомогательной информацией
		AuxFields() []DataFieldDefinition

		// Поля являющиеся ссылками на другие таблицы
		RelationFields() []DataFieldDefinition

		// Значения полей таблицы
		Values() map[string]DataFieldValue

		// Установка полей таблицы
		SetValues(map[string]DataFieldValue) error
	}
)

const (
	FieldTypeUnknown FieldType = iota
	FieldTypeString
	FieldTypeBigint
	FieldTypeInt
	FieldTypeFloat
	FieldTypeBool
	FieldTypeEnum
	FieldTypeDate
)

func NewDataTableDefinion(name string, title string, fields []DataFieldDefinition) *DataTableDefinion {
	return &DataTableDefinion{
		Name:   name,
		Title:  title,
		Fields: fields,
	}
}
