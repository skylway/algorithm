package main

import "fmt"

// 两链表相加
// 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

// 示例：

// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
// 输出：7 -> 0 -> 8
// 原因：342 + 465 = 807

type Node struct {
	val int
	next *Node
}

func (n *Node) SetVal(v int) {
	n.val = v
}

func (n *Node) SetNext(t *Node) {
	n.next = t
}


func AddNode(n1, n2 *Node) (n3 *Node) {
	addition := 0
	var self *Node
	for ;n1 != nil || n2 != nil || addition != 0; n1, n2 = n1.next, n2.next {
		v := n1.val + n2.val + addition
		if v >= 10 {
			v = 10 - v
			addition = 1
		} else {
			addition = 0
		}
		tmp := new(Node)
		tmp.SetVal(v)
		fmt.Printf("%v %p\n", tmp, &tmp)
		
		if n3 == nil {
			n3 = tmp
		} else {
			self.SetNext(tmp)
		}
		self = tmp
		fmt.Printf("%v %p\n", n3, &n3)
	}

	return
}

func main() {
	l1 := new(Node)
	l2 := new(Node)
	l3 := new(Node)
	l4 := new(Node)
	l5 := new(Node)
	l6 := new(Node)
	l1.SetVal(2)
	l2.SetVal(4)
	l3.SetVal(3)
	l4.SetVal(5)
	l5.SetVal(6)
	l6.SetVal(4)

	l1.SetNext(l2)
	l2.SetNext(l3)
	l4.SetNext(l5)
	l5.SetNext(l6)

	// fmt.Printf("%v %p\n", l1, &l1)
	// fmt.Printf("%v %p\n", l2, &l2)
	// fmt.Printf("%v %p\n", l3, &l3)
	// fmt.Printf("%v %p\n", l4, &l4)
	// fmt.Printf("%v %p\n", l5, &l5)
	// fmt.Printf("%v %p\n", l6, &l6)

	to := AddNode(l1,l4)
	fmt.Printf("%v %p\n", to, &to)

}

