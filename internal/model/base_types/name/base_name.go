package base_name

type Value string

func (x Value) Basis() string {
	return string(x)
}
