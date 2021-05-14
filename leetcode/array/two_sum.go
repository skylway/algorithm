package main

import "fmt"

//在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标。
// Given nums = [2, 7, 11, 15], target = 9,
// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1]

// 这道题最优的做法时间复杂度是 O(n)。

// 顺序扫描数组，对每一个元素，在 map 中找能组合给定值的另一半数字，如果找到了，直接返回 2 个数字的下标即可。
// 如果找不到，就把这个数字存入 map 中，等待扫到“另一半”数字的时候，再取出来返回结果。


func addSum(num []int, target int) (result []int) {
	length := len(num)
	if length <= 0 {
		return 
	}
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if num[i] + num[j] == target {
				result = []int{i, j}
				return
			}
		}
	}
	return
}

func addSum2(num []int, target int) (result []int){
	tmpArr := make(map[int]int)
	for k, v := range num {
		if idx, ok := tmpArr[target - v]; ok {
			return []int{idx, k}
		}
		tmpArr[v] = k
	}
	return
}


func main() {
	num := []int{2,3,7,15}
	target := 9
	fmt.Println(addSum2(num, target))
}