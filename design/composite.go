package design

import "fmt"

// 组合模式统一对象和对象集，使得使用相同接口使用对象和对象集。

// 组合模式常用于树状结构，用于统一叶子节点和树节点的访问，并且可以用于应用某一操作到所有子节点。


type ComponentN interface {
	Parent() ComponentN
	SetParent(ComponentN)
	Name() string
	SetName(string)
	AddChild(ComponentN)
	Print(string)
}

const (
	LeafNode = iota
	CompositeNode
)

func NewComponentN(kind int, name string) ComponentN {
	var c ComponentN
	switch kind {
	case LeafNode:
		c = NewLeaf()
	case CompositeNode:
		c = NewComposite()
	}
	c.SetName(name)
	return c
}

type component struct {
	parent ComponentN
	name  string
}

func (c *component) Parent() ComponentN {
	return c.parent
}

func (c *component) SetParent(parent ComponentN) {
	c.parent = parent
}

func (c *component) Name() string {
	return c.name
}

func (c *component) SetName(name string) {
	c.name = name
}

func (c *component) AddChild(ComponentN) {}

func (c *component) Print(string) {}

type Leaf struct {
	component
}

func NewLeaf() *Leaf {
	return &Leaf{}
}

func (c *Leaf) Print(pre string) {
	fmt.Printf("%s--%s\n", pre, c.Name())
}

type Composite struct {
	component
	childs []ComponentN
}

func NewComposite() *Composite {
	return &Composite{
		childs: make([]ComponentN, 0),
	}
}

func (c *Composite) AddChild(child ComponentN) {
	child.SetParent(c)
	c.childs = append(c.childs, child)
}

func (c *Composite) Print(pre string) {
	fmt.Printf("%s+%s\n", pre, c.Name())
	pre += " "
	for _, comp := range c.childs {
		comp.Print(pre)
	}
}