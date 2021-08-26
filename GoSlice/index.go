package GoSlice

import (
	"fmt"
	"reflect"
)

/**
 * Golang中的切片等同于Java的列表，即不定长数组
 * 切片可以通过数组初始化，也可以通过如下两种方式初始化
 */

var StrSlice1 []string

var StrSlice2 = make([]string, 3)
//var StrSlice2 = make([]string, 0)

/**
 * Golang中任何类型使用interface{}表示，类似于Java中的顶级父类Object
 */

var AnySlice []interface{}

var AS = make([]interface{}, 0)

func PrintLenAndCap(slice []string) {
	fmt.Println(len(slice), cap(slice))
}

func ToString(slice []string) string {
	var res = "["
	l := len(slice)
	for i:=0; i< l; i++ {
		if i == l - 1 {
			res += slice[i]
		} else {
			res += slice[i] + ", "
		}
	}
	res += "]"
	return res
}

func Test() {
	fmt.Println(reflect.TypeOf(StrSlice1))

	PrintLenAndCap(StrSlice1)
	fmt.Println(ToString(StrSlice1))

	StrSlice1 = append(StrSlice1, "asdf")
	PrintLenAndCap(StrSlice1)
	fmt.Println(ToString(StrSlice1))

	StrSlice1 = append(StrSlice1, "qwer")
	PrintLenAndCap(StrSlice1)
	fmt.Println(ToString(StrSlice1))

	fmt.Println("==============================================")

	PrintLenAndCap(StrSlice2)
	fmt.Println(ToString(StrSlice2))

	StrSlice2 = append(StrSlice2, "123")
	PrintLenAndCap(StrSlice2)
	fmt.Println(ToString(StrSlice2))
}

func TestAnySlice() {
	AnySlice = append(AnySlice, "aaa")
	AnySlice = append(AnySlice, 123)
	AnySlice = append(AnySlice, []int{1, 5, 6, 7, 9})
}

