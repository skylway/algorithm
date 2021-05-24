package design

import "testing"


func TestOperator(t *testing.T) {
	var factory OperatorFactory
	factory = PlusOperatorFactory{}
	if compute(factory, 1, 2) != 3 {
		t.Fatal("error with factory method pattern")
	}
	factory = MinusOperatorFactory{}
	if compute(factory, 5, 2) != 3 {
		t.Fatal("error with factory method pattern")
	}
}

func compute(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}