package design

import (
	"testing"
)

func TestNewComposite(t *testing.T) {
	// tests := []struct {
	// 	name string
	// 	want *Composite
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if got := NewComposite(); !reflect.DeepEqual(got, tt.want) {
	// 			t.Errorf("NewComposite() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
	root := NewComponentN(CompositeNode, "root")
	c1 := NewComponentN(CompositeNode, "c1")
	c2 := NewComponentN(CompositeNode, "c2")
	c3 := NewComponentN(CompositeNode, "c3")

	l1 := NewComponentN(LeafNode, "l1")
	l2 := NewComponentN(LeafNode, "l2")
	l3 := NewComponentN(LeafNode, "l3")

	root.AddChild(c1)
	root.AddChild(c2)
	c1.AddChild(c3)
	c1.AddChild(l1)
	c2.AddChild(l2)
	c2.AddChild(l3)
	root.Print("")
}
