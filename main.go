package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	//fmt.Println("Hello World!")

	/*fmt.Print("\n====================================================================\n")
	printSomething()
	fmt.Print("====================================================================\n\n")*/

	//person.GoPrint()
	//fmt.Println(person.ToString())
	//person.ForPrintNum()

	//array.Run()

	//slice.Test()

}

func printSomething() {
	var a = "asdf"
	var b = 1
	c := "qwer" // go 变量声明定义 语法糖写法
	d := 123.45
	res := float64(b) + d
	fmt.Println(a, b, c, d, res)
	fmt.Println(reflect.TypeOf(c), reflect.TypeOf(b), reflect.TypeOf(c), reflect.TypeOf(d))

	var pi = 3.1415926
	fmt.Println(strconv.FormatFloat(pi, 'f',9,64))
}

