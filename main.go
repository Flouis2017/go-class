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

	/*Person.GoPrint()
	fmt.Println(Person.ToString())
	Person.ForPrintNum()*/

	/*var a = GoArray.Zoo[2:]
	fmt.Println(reflect.TypeOf(a), ">>", a)
	fmt.Println(reflect.TypeOf(GoArray.Zoo), ">>", GoArray.Zoo)
	GoArray.Run()*/

	//GoSlice.Test()
	/*fmt.Println(reflect.TypeOf(GoSlice.StrSlice1), reflect.TypeOf(GoSlice.StrSlice2), reflect.TypeOf(GoSlice.StrSlice3))
	fmt.Println("GoSlice.StrSlice1: ", GoSlice.StrSlice1)
	fmt.Println("GoSlice.StrSlice2: ", GoSlice.StrSlice2)
	GoSlice.TestAnySlice()
	fmt.Println(GoSlice.AnySlice)*/

	/*GoMap.Init()
	fmt.Println(reflect.TypeOf(GoMap.MapOne))
	fmt.Println("GoMap.MapOne: ", GoMap.MapOne)
	fmt.Println("GoMap.MapTwo: ", GoMap.MapTwo)
	fmt.Println(GoMap.MapOne["city"], GoMap.MapOne["name"] == "")
	// 遍历 map
	for k,v := range GoMap.MapOne {
		fmt.Println(k, v)
	}
	// map 转 json字符串
	marshal, _ := json.Marshal(GoMap.MapOne)
	jsonStr := string(marshal)
	fmt.Println(jsonStr)

	fmt.Println(GoMap.AnyMap, len(GoMap.AnyMap))
	delete(GoMap.AnyMap, "arr")
	fmt.Println(GoMap.AnyMap, len(GoMap.AnyMap))*/

	//GoPointer.IntPointerTest()
	//GoPointer.ListPointerTest()
	//GoPointer.PointFuncTest()

	/*GoStruct.FirstTest()
	GoStruct.SecondTest()
	GoStruct.ThirdTest()
	ref := GoStruct.FourthTest()
	fmt.Println(*ref)
	fmt.Println(GoStruct.MyStudent)

	GoStruct.AnimalTest()*/
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

