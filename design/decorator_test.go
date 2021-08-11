package design

import (
	"testing"
)

func TestDecorator(t *testing.T) {
	var c Component = &ConcreteComponent{}
	c = WarpAddDecorator(c, 10)
	c = WarpMulDecorator(c, 8)
	res := c.Calc()
	t.Logf("res %d\n", res)
}
