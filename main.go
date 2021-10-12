package main

import (
	"errors"
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"time"
	"unsafe"
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
	test := 16132 % 64
	fmt.Println(test)
	var musicStatus uint32
	status := "1"
	ttt, err := strconv.ParseUint(status, 0, 32)
	musicStatus = uint32(ttt)
	fmt.Println(musicStatus)
	switch status {
	case "0":
		fmt.Println(0)
	case "1":
		fmt.Println(1)
		fmt.Println(111111)
	case "2":
		fmt.Println(2)
	}

	type testStruct struct {
		Musint string
		Id  int
	}
	data := make(map[int]*testStruct)
	// data = map[int]string {
	// 	1: "1",
	// 	2: "2",
	// }
	data[3] = &testStruct{Musint: "test1", Id: 3}

	fmt.Println(data[2])
	overflow(2147483647)
	a := 10
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a)
	fmt.Printf("b:%d type:%T ptr1: %p ptr2 %p\n", *b,b,b,&b)
	// var test sync.Map
	f := testFunc()
	f()
	tmp1 := add(10)
	fmt.Println(tmp1(1), tmp1(2))
	tmp2 := add(100)
	fmt.Println(tmp2(1), tmp2(2))
	testDefer()
	testFoo(2, 0)
	list := [6]int{1,2,3,4,5,6}
	slice := list[:]
	slice2 := slice
	fmt.Println(slice, unsafe.Pointer(&slice), slice2, unsafe.Pointer(&slice2), list)
	// type Student struct {
	// 	name string
	// }
	// m := map[string]*Student{"people": {"zhoujielun"}}
	// m["people"].name = "wuyanzu";
	// fmt.Println(m);
	// ret := exec("111", func(n string) string {
	// 	return n + "func1"
	// }, func(n string) string {
	// 	return n + "func2"
	// }, func(n string) string {
	// 	return n + "func3"
	// }, func(n string) string {
	// 	return n + "func4"
	// })
	// fmt.Println(ret)
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	// for i := 0; i < 10; i++ {
	// 	go func(i int) {
	// 		fmt.Println("i: ", i)
	// 		wg.Done()
	// 	}(i)
	// }
	wg.Wait()
}

//删除函数
func remove(s []string, i int) []string {
    return append(s[:i], s[i+1:]...)
}

func overflow(numControlByUser int32) {
	var numInt int32 = 0
	numInt = numControlByUser + 1
	// 对长度限制不当，导致整数溢出
	// fmt.Printf("%d\n", numInt)
	// 使用numInt，可能导致其他错误
	if numInt < 0 {
		fmt.Println("integer overflow")
		return
	}
	fmt.Println("integer ok")	
}

func testFunc() func() {
	x := 500
	fmt.Printf("x (%p) = %d\n", &x, x)
	return func(){
		fmt.Printf("x (%p) = %d\n", &x, x)
	}
}

func add(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

func fibonaci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonaci(i-1) + fibonaci(i-2)
}

func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}

func testDefer() {
	var whatever [5]struct{}
	for i := range whatever {
		defer func(x int) {fmt.Println(x)} (i)
	}
}

func testFoo(a, b int) (i int, err error) {
	defer fmt.Printf("first defer err %v\n", err)
	defer func(err error) {fmt.Printf("second defer err %v\n", err)} (err)
	defer func() { fmt.Printf("third defer err %v\n", err) } ()
	if b == 0 {
		err = errors.New("divided by zero!")
		return
	}
	i = a / b
	return
}
type query func(string) string

func exec(name string, vs ...query) string {
	ch := make(chan string)
	fn := func(i int) {
		fmt.Println(vs[i])
		ch <- vs[i](name)
	}
	for i, _ := range vs {
		go fn(i)
	}
	return <-ch
}