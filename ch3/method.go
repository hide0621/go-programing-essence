package ch3

type Value int

func (v Value) Add1(n Value) Value {
	return v + n
}
