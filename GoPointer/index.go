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
 * 数组(切片)指针(*[]<type>)和指针数组(切片)([]*<type>)，两者不同，不要混淆了！
 * 加一个定语就会比较好理解：数组指针 = 数组的指针，*ref之后是一个数组。
 * 而指针数组是一个存放指针的数组
 * 数组和切片同理，下面以切片做示例
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
	l := len(list)
	for i := 0; i < l; i++ {
		refSlice = append(refSlice, &list[i])
	}
	ll := len(refSlice)
	for i := 0; i < ll; i++ {
		fmt.Print(*refSlice[i], "\t")
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

