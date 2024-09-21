package basis

import (
	"time"
)

type (
	SingleBased[T any] struct {
		Val T
	}

	ID struct {
		SingleBased[int64]
	}
	Name struct {
		SingleBased[string]
	}
	Flag struct {
		SingleBased[bool]
	}
	Qty struct {
		SingleBased[float64]
	}
	Date struct {
		SingleBased[time.Time]
	}
	Moment struct {
		SingleBased[time.Time]
	}
	Amount struct {
		SingleBased[float64]
	}
	Phone struct {
		SingleBased[string]
	}

	TableRecord[IDFields any, BaseFields any, ExtFields any, RefFields any] struct {
		ID   IDFields
		Base BaseFields
		Ext  ExtFields
		Refs RefFields
	}

	NewValuer[T any] interface {
		NewValue() T
	}
)

func (sb SingleBased[T]) Base() T {
	return sb.Val
}

func (t *TableRecord[IDFields, BaseFields, ExtFields, RefFields]) SetID(id IDFields) {
	t.ID = id
}

func (t *TableRecord[IDFields, BaseFields, ExtFields, RefFields]) GetID() IDFields {
	return t.ID
}
