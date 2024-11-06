package common

import "time"

type (
	ID int64

	Name string
	Date time.Time

	Option[T any] struct {
		Setted bool // Is value setted and valid
		Value  T
	}

	Result[T any] struct {
		Value T
		Err   error
	}
)
