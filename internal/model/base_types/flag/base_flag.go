package base_flag

type Value bool

func (x Value) Basis() bool {
	return bool(x)
}
