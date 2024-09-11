package base_date

import "time"

type Value time.Time

func (x Value) Basis() time.Time {
	return time.Time(x)
}
