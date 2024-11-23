package common

type (
	ID int64

	Name  string
	Notes string

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

func CollectIds(src map[int]TableField, filter func(fld TableField) (canAdd bool)) []int {
	result := make([]int, 0)
	for k, v := range src {
		if filter(v) {
			result = append(result, k)
		}
	}
	return result
}

func CollectNames(src map[int]TableField, ids []int) []string {
	result := make([]string, 0, len(ids))
	for _, v := range ids {
		tf, ok := src[v]
		if ok {
			result = append(result, tf.Name)
		} else {
			result = append(result, "null")
		}
	}
	return result
}
