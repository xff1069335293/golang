package main

import (
	"fmt"
	//"time"
) //导入包
/*func main()  {
	fmt.Println("hello")
	fmt.Printf("你好\n")
	time.Sleep(time.Second)
}*/
/*const pi float32  = 3.14
var num int  = 55
func get_val()(int ,int)  {
	return 4,5
}
const (
	n1,n2 = iota+1,iota+2//1 2 2 3
	n3,n4 = iota +1 ,iota +2
)*/
const (
	kb = 1 << (10*(iota +1))
	mb
	gb
)
func main()  {
	/*var i int = 99
	a := 88//类型推导 var a int = 88 只允许局部变量
	b := 77
	var c,d int = 3, 4
	//var name string = "小王"
	//var age int  = 44
	var name,age = "小李" ,24
	fmt.Printf("type:= %T name = %s,type:%T age = %d\n",name ,name ,age ,age)
	fmt.Println(name,age)
	fmt.Println(i,a,b,c,d,num)
	a , b := 2,3
	a, b = b,a
	fmt.Println(a,b)
	n,m := get_val()
	fmt.Println(n1,n2,n3,n4)
	fmt.Println(kb,mb,gb)
	var a = 3
	var b = 0b101
	fmt.Println(b)
	fmt.Printf("%d\n%b\n",a,a)//3 11
	var b bool = false
	b2 := true
	fmt.Println(b,b2)*/
	var s string 
	s = "hello"
	fmt.Println(s)
}