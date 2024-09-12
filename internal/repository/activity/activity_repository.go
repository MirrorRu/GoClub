package activity

import (
	"goclub/internal/model"
	"goclub/internal/repository"
)

type (
	ActivityFilter struct {
		nameFilter string
	}

	ActivityStorage repository.TableStorage[model.ActivityID, model.Activity, ActivityFilter]
)

/*
	TableStorage[IDType any, TableStructType any, FilteringType any] interface {
		Create(TableStructType) (IDType, error)
		Read(IDType) (TableStructType, error)
		Update(TableStructType) error
		Delete(IDType) error
		List(FilteringType, ListMakeOption) ([]TableStructType, error)
	}
*/
