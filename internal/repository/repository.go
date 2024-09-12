package repository

import "goclub/internal/model/basis"

type (
	StringFilter string

	ListMakeOption uint32

	TableStructTypeDef interface {
		basis.Table[any, any, any, any]
	}

	TableStorage[IDType any, TableStructType any, FilteringType any] interface {
		Create(TableStructType) (IDType, error)
		Read(IDType) (TableStructType, error)
		Update(TableStructType) error
		Delete(IDType) error
		List(FilteringType, ListMakeOption) ([]TableStructType, error)
	}
)

const (
	ListMakeAll ListMakeOption = 1 << iota
	ListMakeWithIDs
	ListMakeWithBase
	ListMakeWithExt
	ListMakeWithRefs
)
