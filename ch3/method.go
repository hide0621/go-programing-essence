package ch3

type Value int

func (v Value) Add(n Value) Value {
	return v + n
}
