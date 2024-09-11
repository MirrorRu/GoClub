package base_id

type Value int

func (x Value) Basis() int {
	return int(x)
}
