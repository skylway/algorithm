package main

import (
	"algorithm/leetcode/libary"
	"fmt"
)

//中根遍历一颗树。
// Example :

// ```c
// Input: [1,null,2,3]
//    1
//     \
//      2
//     /
//    3

// Output: [1,3,2]
// ```


type TreeNode = libary.TreeNode

func InOrderTraversal(root *TreeNode, result *[]int) {
	if root != nil {
		InOrderTraversal(root.Left, result)
		*result = append(*result, root.Val)
		InOrderTraversal(root.Right, result)
	}
	return
}

func main() {
	node := Ints2TreeNode([]int{2, 4, 2, 4, 6, 5})
	var result []int
	InOrderTraversal(node, &result)
	fmt.Println(result)
}

var NULL = -1 << 63

// Ints2TreeNode 利用 []int 生成 *TreeNode
func Ints2TreeNode(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: ints[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && ints[i] != NULL {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && ints[i] != NULL {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
