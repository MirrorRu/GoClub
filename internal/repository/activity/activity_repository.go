package activity

type (
	ActivityFilter struct {
		nameFilter string
	}

	//ActivityStorage repository.TableStorager[model.ActivityID, model.Activity, ActivityFilter]
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
