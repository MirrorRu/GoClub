package base_phone

type Value string

func (x Value) Basis() string {
	return string(x)
}
