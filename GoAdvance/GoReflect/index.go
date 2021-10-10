package GoReflect

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age int
	Gender int
}

type Student struct {
	User
	Class string
}

func ReflectTest(arg interface{}) {
	var t = reflect.TypeOf(arg)
	fmt.Println("反射获取类型：", t)
	fmt.Println("反射获取种类：", t.Kind())

	var v = reflect.ValueOf(arg)
	fmt.Println("反射获取数据：", v)

	var n = t.NumField()
	fmt.Println("反射获取属性个数：", n)

	fmt.Print("反射按属性顺序取值：")
	for i := 0; i < n; i++ {
		fmt.Print(v.Field(i), "\t")
	}
}

func ReflectTest2(arg interface{}) {
	var t = reflect.TypeOf(arg)
	fmt.Println("反射获取类型：", t)

	var v = reflect.ValueOf(arg)
	fmt.Println("反射获取数据：", v)

	fmt.Println(v.FieldByIndex([]int{0}))
	fmt.Println(v.FieldByIndex([]int{0, 1}))
	fmt.Println(v.FieldByName("Class"))
}

func ReflectModify(arg interface{}) {
	var e = reflect.ValueOf(arg).Elem()
	e.FieldByName("Class").SetString("四年二班")
}
