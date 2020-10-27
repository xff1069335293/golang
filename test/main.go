package main

import "fmt"

func main() {

	//变量的声明
	// var i int //整型
	// fmt.Println(i)
	// var s string//字符串类型
	// fmt.Println(s)
	// var pi *int//指针类型
	// fmt.Println(pi)
	// var stu struct {//结构体类型
	// 	id   int
	// 	name string
	// }
	// fmt.Println(stu.id, stu.name)

	//for遍历
	// var arr [10]int//整型数组类型
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Printf("%d ", arr[i])
	// }

	//range遍历
	//var a1 [5]int = [5]int{1, 2, 3, 4, 5}
	//var a2 = [5]int{2, 3, 4, 5, 6}
	// a3 := [5]int{3, 4, 5, 6, 7}
	// var vi []int = a3[:2]
	v1 := make([]int, 5, 10)
	v1 = []int{1, 2, 3, 4, 5, 6}
	for k, v := range v1 {
		fmt.Printf("%d->%d ", k, v)
	}
	fmt.Printf("%d %d", len(v1), cap(v1))
	// var mp map[string]int //map
	// var fun func(i int) int//函数

	// var (
	// 	i1 int
	// 	s1 string
	// )

	// //变量初始化
	// var v1 int = 10
	// var v2 = 20
	// v3 := 30
}
