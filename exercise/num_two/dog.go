package main

import (
	"fmt"
	"math"
)

// 说汪星人的智商能达到人类 4 岁儿童的水平，更有些聪明汪会做加法计算。比如你在地上放两堆小球，分别有 1 只球和 2 只球，聪明汪就会用“汪！汪！汪！”表示 1 加 2 的结果是 3。
// 本题要求你为电子宠物汪做一个模拟程序，根据电子眼识别出的两堆小球的个数，计算出和，并且用汪星人的叫声给出答案。
// 输入格式:
// 输入在一行中给出两个 [1, 9] 区间内的正整数 A 和 B，用空格分隔。
// 输出格式:
// 在一行中输出 A + B 个Wang!。

// func main() {
// 	var m,n int
// 	var total int
// 	fmt.Scan(&m, &n)
// 	total = m + n
// 	s := ""
// 	for i := 0; i < total; i++ {
// 		s += "Wang!"
// 	}
// 	fmt.Print(s)
// }

// 给定一个完全由小写英文字母组成的字符串等差递增序列，该序列中的每个字符串的长度固定为 L，从 L 个 a 开始，以 1 为步长递增。例如当 L 为 3 时，序列为 { aaa, aab, aac, ..., aaz, aba, abb, ..., abz, ..., zzz }。这个序列的倒数第27个字符串就是 zyz。对于任意给定的 L，本题要求你给出对应序列倒数第 N 个字符串。

// 输入格式:
// 输入在一行中给出两个正整数 L（2 ≤ L ≤ 6）和 N（≤10

// 输出格式:
// 在一行中输出对应序列倒数第 N 个字符串。题目保证这个字符串是存在的。

// 输入样例1:
// 3 7417
// 结尾无空行

// 输出样例1:
// pat

// 结尾无空行

// 思路:
// 这题看上去很唬人，其实只是一个26进值的转化而已。
// 首先知道总数是多少，才能对各个进行位置进行求余操作。
func main() {

	var L,N int
	fmt.Scan(&L, &N)
	a := make([]rune, 7)
	sum := math.Pow(26, float64(L)) //求出总数
	tmp := int(sum) - N //然后计算出一共要多少位，因为是倒数第N位，减去就是这个了正数第几位了
	fmt.Printf("total: %d\n", int(sum))
	fmt.Printf("tmp: %d\n", tmp)
	//然后用这个数对26不断求余，除整就够了。
	for i := L - 1; i >= 0; i-- {
		t := tmp%26
		tt := rune(t)
		fmt.Printf("t: %d\n", t)
		fmt.Printf("tt: %d\n", t)
		a[i] = 'a' + tt
		tmp /= 26
		fmt.Printf("ttmp: %d\n", tmp)
	}
	for i := 0; i < L; i++ {
		fmt.Printf("%s", string(a[i]))
	}

}