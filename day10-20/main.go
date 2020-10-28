package main

import "fmt"

func main() {
	//string.Split分割函数  len
	//var s1 string = "hello"
	// s2 := "how are you"
	// s3 := strings.Split(s2, " ")
	// fmt.Println(s3[1])

	// s := "hello 你好"
	// for i := 0; i < len(s); i++ { //按字节计算
	// 	fmt.Printf("%c\n", s[i])
	// }
	// for i, v := range s { //按单个字符输出
	// 	fmt.Printf("index = %d,v=%c\n", i, v)
	// }
	// var c1 byte = 'a'
	// var c2 rune = '中'
	// fmt.Printf("%T %T\n", c1, c2)
	// fmt.Println(c1, c2)

	//if
	// if b := 3; b == 3 {
	// 	fmt.Println("==")
	// }

	//for range
	// s := "hello 你好"
	// for _, v := range s {
	// 	fmt.Println(v)
	// }

	//switch
	// score := 'A'
	// switch score {
	// case 'A':
	// 	fmt.Println("成绩为:90-100")
	// case 'B':
	// 	fmt.Println("成绩为:80-90")
	// }

	//var arr [5]int=[5]int{1,2,3,4,5}
	//arr := [5]int{1, 2, 3, 4, 5}
	//var arr = [5]int{1, 2, 3, 4, 5}
	//var arr [5]int = [5]int{1: 3, 4: 9}
	//arr := [...]int{1, 2, 4}
	//fmt.Println(arr)
	//fmt.Printf("%T", arr)

	// var a = [3][2]string{{"陕西", "西安"}, {"四川，“成都"}, {"中国", "北京"}}
	// fmt.Println(a)
	// var b [3][2]string
	// b = a
	// a[0][1] = "长安"
	// fmt.Println(b)
	// fmt.Println(a)

	var a []int = []int{1, 2, 3}
	if a == nil {
		fmt.Println("没有值")
	}
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println("len=", len(a), "cap =", cap(a))
}
