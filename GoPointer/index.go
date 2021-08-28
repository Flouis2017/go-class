package GoPointer

import (
	"fmt"
	"reflect"
)

func IntPointerTest() {
	var a int
	a = 100
	fmt.Println(a)

	var b *int // 声明指针
	b = &a // 取地址
	*b = 999
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(a, b)
	fmt.Println(a == *b, &a == b)
}


/**
 * 数组指针(*[]<type>)和指针数组([]*<type>)，两者不同，不要混淆了！
 * 加一个定于就比较好理解：数组指针 = 数组的指针，*ref之后是一个数组。
 * 而指针数组是一个存放指针的数组
 * 数组和切片同理，下面以切面做示例
 */

func ListPointerTest() {
	var list = []string{"aa", "bb", "cc", "dd", "ee"}
	fmt.Println(list)

	var ref *[]string
	ref = &list
	fmt.Println(*ref)

	(*ref)[2] = "xx"
	fmt.Println(list, *ref)

	fmt.Println("================================")

	var refSlice []*string
	var str1 = "str1"
	var str2 = "str2"
	var str3 = "str3"
	var str4 = "str4"
	var str5 = "str5"
	refSlice = append(refSlice, &str1)
	refSlice = append(refSlice, &str2)
	refSlice = append(refSlice, &str3)
	refSlice = append(refSlice, &str4)
	refSlice = append(refSlice, &str5)
	for i := 0; i < len(refSlice); i++ {
		fmt.Print(*refSlice[i], " & ")
	}
}

func PointFuncTest() {
	str := "xxx"
	fmt.Println("Before func change: ", str)
	change(&str)
	fmt.Println("After func change: ", str)
}

func change(p *string) {
	*p = "Golang to be No.1"
}

