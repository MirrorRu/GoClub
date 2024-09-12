package basis

import "time"

type BaseSingle[T any] struct {
	Val T
}

func (t BaseSingle[T]) Basis() T {
	return t.Val
}

type (
	ID     BaseSingle[int64]
	Name   BaseSingle[string]
	Flag   BaseSingle[bool]
	Qty    BaseSingle[float64]
	Date   BaseSingle[time.Time]
	Moment BaseSingle[time.Time]
	Amount BaseSingle[float64]
	Phone  BaseSingle[string]
)

type StructType interface{}

type Table[IDFields any, BaseFields any, ExtFields any, RefFields any] struct {
	ID   IDFields
	Base BaseFields
	Ext  ExtFields
	Refs RefFields
}
