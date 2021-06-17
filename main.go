package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func main() {
// 	fsm := fsm.NewFSM(
// 		"idle",
// 		fsm.Events{
// 			{Name: "scan", Src: []string{"idle"}, Dst: "scanning"},
// 			{Name: "working", Src: []string{"scanning"}, Dst: "scanning"},
// 			{Name: "situation", Src: []string{"scanning"}, Dst: "scanning"},
// 			{Name: "situation", Src: []string{"idle"}, Dst: "idle"},
// 			{Name: "finish", Src: []string{"scanning"}, Dst: "idle"},
// 		},
// 		fsm.Callbacks{
// 			"scan": func(e *fsm.Event) {
// 				fmt.Println("after_scan: " + e.FSM.Current())
// 			},
// 			"working": func(e *fsm.Event) {
// 				fmt.Println("working: " + e.FSM.Current())
// 			},
// 			"situation": func(e *fsm.Event) {
// 				fmt.Println("situation: " + e.FSM.Current())
// 			},
// 			"finish": func(e *fsm.Event) {
// 				fmt.Println("finish: " + e.FSM.Current())
// 			},
// 		},
// 	)

// 	fmt.Println(fsm.Current())

// 	err := fsm.Event("scan")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println("1:" + fsm.Current())

// 	err = fsm.Event("working")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println("2:" + fsm.Current())

// 	err = fsm.Event("situation")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println("3:" + fsm.Current())

// 	err = fsm.Event("finish")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println("4:" + fsm.Current())

// }


func main() {
	letter := `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	bytes := []byte(letter)
	rand.Seed(time.Now().UnixNano())
	fmt.Println(len(bytes))
	result := []byte{}

	num := rand.Intn(len(bytes))
	result = append(result, bytes[num])

	fmt.Println(string(result))

	t := time.Now()
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err == nil {
		t = t.In(loc)
	}
	//t = t.Format("2006-01-02 15:04:05")
	fmt.Printf("当前的时间是: %d-%d-%d %d:%d:%d\n", t.Year(),
		t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()) 
		

	//时间转换的模板，golang里面只能是 "2006-01-02 15:04:05" （go的诞生时间）
	timeTemplate1 := "2006-01-02 15:04:05" //常规类型
	// timeTemplate2 := "2006/01/02 15:04:05" //其他类型
	timeTemplate3 := "2006-01-02"          //其他类型
	// timeTemplate4 := "15:04:05"            //其他类型
	
	//当前日期（年-月-日）的时间戳
	todaystr := time.Now().Format("2006-01-02")
	//使用parseInLocation将字符串格式化返回本地时区时间
	todayint, _ := time.ParseInLocation(timeTemplate3, todaystr, loc) 
	fmt.Println("日期：", todaystr, "的时间戳是:", todayint.Unix()) 
	//输出：日期： 2021-04-23 的时间戳是: 1619107200
	
	//当前时间（年-月-日 时：分：秒）的时间戳
	todaystr = time.Now().Format("2006-01-02 15:04:05")
	todayint, _ = time.ParseInLocation(timeTemplate1 , todaystr, loc) 
	fmt.Println("日期：", todaystr, "的时间戳是:", todayint.Unix())
	//输出：日期： 2021-04-23 15:44:19 的时间戳是: 1619163859
	fmt.Printf("当前的时间是: %d-%d-%d %d:%d:%d\n", todayint.Year(),
		todayint.Month(), todayint.Day(), todayint.Hour(), todayint.Minute(), todayint.Second()) 

    s := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
    fmt.Println(s)
    s = append(s[:0], s[2:]...)
    s = append(s[:6], s[7:]...)
	fmt.Println(s)
	ta := ``
	fmt.Println(len(ta))
}

//删除函数
func remove(s []string, i int) []string {
    return append(s[:i], s[i+1:]...)
}