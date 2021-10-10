package GoPointer

import (
	"fmt"
	"reflect"
)

/**
 * var x *<类型>：声明x为某个类型的指针，指针是特殊的类型，专门存放变量的内存地址
 * x = &<变量>：获取某个变量的内存地址，并将其赋值给x。该过程也叫指针的定义，可以和上面的声明合起来写，即var x *<类型> = &<变量> —— 声明定义
 * *x：通过指针操作变量
 * 例如：
 * var a = "Flower" // 声明定义一个字符串变量a，值为“Flower”
 * var aRef *string = &a // 声明定义一个字符串指针aRef用于存放变量a的内存地址
 * *aRef // 通过指针访问指针所指向的变量，*aRef修改后，aRef指向的a也会相应被修改
 */

func StringPointerTest() {
	var str = "abc"
	var a = "Flower"
	var aRef *string = &a
	fmt.Println(*aRef)
	*aRef = str
	fmt.Println(a)
}

func IntPointerTest() {
	var a int
	a = 100
	fmt.Println(a)

	var b *int // 声明指针
	b = &a // 获取地址
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

// ListChange1 值传递能改变到原变量
func ListChange1(arg []string) {
	arg = append(arg, "zzz")
}

// ListChange2
/**
 * 指针传递能修改到原变量，在Java中也叫引用传递，Java的引用等同于C(++)/Golang中的指针，只是Java将指针操作（&/*）在语言层面上屏蔽了
 * 不同教程对这两种传递还有不能的命名：我上大学时，C(++)老师把指针/引用传递称为“实参传递”，把值传递称为“形参传递”。
 */
func ListChange2(p *[]string) {
	*p = append(*p, "zzz")
}

