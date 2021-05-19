package main

import (
	"algorithm/leetcode/libary"
	"fmt"
)

// ## 题目大意

// 给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。

// ## 解题思路

// - 输出 1~n 元素组成的 BST 所有解。这一题递归求解即可。外层循环遍历 1~n 所有结点，作为根结点，内层双层递归分别求出左子树和右子树。


type TreeNode = libary.TreeNode


func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generateBSTree(1, n)
}

func generateBSTree(start, end int) []*TreeNode {
	tree := []*TreeNode{}
	if start > end {
		tree = append(tree, nil)
		return tree
	}
	for i := start; i <= end; i++ {
		left := generateBSTree(start, i-1)
		right := generateBSTree(i+1, end)
		for _, l := range left {
			for _, r := range right {
				root := &TreeNode{Val: i, Left: l, Right: r}
				tree = append(tree, root)
			}
		}
	}
	return tree
}

func main() {
	fmt.Println(generateTrees(7))
}