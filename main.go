package main

import (
	"fmt"
	"reflect"
	"strconv"
)

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

func main() {
	//fmt.Println("Hello World!")

	/*fmt.Print("\n====================================================================\n")
	printSomething()
	fmt.Print("====================================================================\n\n")*/

	/*Person.GoPrint()
	fmt.Println(Person.ToString())
	Person.ForPrintNum()*/

	// 数组测试
	/*var a = GoArray.Zoo[2:]
	fmt.Println(reflect.TypeOf(a), ">>", a)
	fmt.Println(reflect.TypeOf(GoArray.Zoo), ">>", GoArray.Zoo)
	GoArray.Run()*/

	// 切片测试
	//GoSlice.Test()
	/*fmt.Println(reflect.TypeOf(GoSlice.StrSlice1), reflect.TypeOf(GoSlice.StrSlice2), reflect.TypeOf(GoSlice.StrSlice3))
	fmt.Println("GoSlice.StrSlice1: ", GoSlice.StrSlice1)
	fmt.Println("GoSlice.StrSlice2: ", GoSlice.StrSlice2)
	GoSlice.TestAnySlice()
	fmt.Println(GoSlice.AnySlice)*/

	// Map测试
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

	// 指针测试
	//GoPointer.IntPointerTest()
	//GoPointer.ListPointerTest()
	//GoPointer.PointFuncTest()
	//var str = []string{"aaa", "bbb", "ccc"}
	//fmt.Println(str)
	//GoPointer.ListChange1(str)
	//fmt.Println(str)
	//GoPointer.ListChange2(&str)
	//fmt.Println(str)

	// 结构体测试
	//GoStruct.FirstTest()
	//var got = GoStruct.SecondTest()
	//fmt.Println(got)
	//GoStruct.ThirdTest()
	//ref := GoStruct.FourthTest()
	//fmt.Println(*ref)
	//fmt.Println(GoStruct.MyStudent)
	//GoStruct.AnimalTest()

	/*================================= golang进阶 =================================*/
	// goroutine测试
	//GoRoutine.ShowTime()
	//GoRoutine.SyncTest()
	//GoRoutine.AsyncTest()
	//GoRoutine.WaitGroupTest()

	// reflect 反射测试
	//var user = GoReflect.User{ Name: "大师兄", Age: 30, Gender: 1 }
	//var student = GoReflect.Student{ User:  user, Class: "三年二班" }
	//GoReflect.ReflectTest(user)
	//GoReflect.ReflectTest2(student)
	//GoReflect.ReflectModify(&student)
	//fmt.Println(student)

	// IO 测试
	//var filePath = "./test.txt"
	//var res = GoIO.ReadTest(filePath)
	//fmt.Println(res)
	//GoIO.WriteTest("./test2.txt")

}

