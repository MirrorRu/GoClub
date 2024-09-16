package basis

import (
	"time"
)

type BaseSingle[T any] struct {
	Val T
}

func (t BaseSingle[T]) Basis() T {
	return t.Val
}

type (
	IDer[T comparable] interface {
		Inc()
		Get() T
	}

	ComparableIDer[T comparable] interface {
		comparable
		IDer[T]
	}

	ID BaseSingle[int64]
)

func (x *ID) Inc() {
	x.Val++
}

func (x ID) Get() ID {
	return x
}

type (
	Name   BaseSingle[string]
	Flag   BaseSingle[bool]
	Qty    BaseSingle[float64]
	Date   BaseSingle[time.Time]
	Moment BaseSingle[time.Time]
	Amount BaseSingle[float64]
	Phone  BaseSingle[string]

	TableRecord[IDFields ComparableIDer[ID], BaseFields any, ExtFields any, RefFields any] struct {
		ID   IDFields
		Base BaseFields
		Ext  ExtFields
		Refs RefFields
	}
)
