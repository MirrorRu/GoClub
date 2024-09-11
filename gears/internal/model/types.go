package model

type ID int64
type Amount float64

type AccrualValue Amount
type WithdrawnValue Amount

type CRUDResult struct {
	RowsAffected int64
}
