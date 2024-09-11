package base_amount

type Value float64

func (x Value) Basis() float64 {
	return float64(x)
}
