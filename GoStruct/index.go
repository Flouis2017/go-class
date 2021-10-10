package GoStruct

import (
	"fmt"
	"reflect"
)

/**
 * Golang不是纯面向对象语言，所以并没有提供class关键字，但是struct，即结构体，可以近似地当作类来使用
 */

type Student struct {
	Name string
	Age int
	Gender string
	Hobbies []string
}

var MyStudent Student

// FirstTest 类似于Java中使用默认无参构造方法创建对象，再使用setter方法初始化成员变量
func FirstTest() {
	var student Student
	student.Name = "Flouis"
	student.Age = 27
	student.Gender = "Male"
	student.Hobbies = []string{"Music", "Dota2"}
	//student.Hobbies = append(student.Hobbies, "Music")
	//student.Hobbies = append(student.Hobbies, "Dota2")
	fmt.Println(student)
}

// SecondTest 类似于Java中使用有参构造方法创建对象
func SecondTest() Student {
	student := Student{
		Name: "Lily",
		Age: 25,
		Gender: "Female",
		Hobbies: []string{"Singing", "Dancing"},
	}
	return student
}

func ThirdTest() {
	student := Student{"xxx", 30, "Unknown", []string{"qwer", "asdf"}}
	fmt.Println(student)
}

// FourthTest 使用new关键字创建对象，这时候需要注意返回的是一个指针
func FourthTest() *Student {
	student := new(Student)
	student.Name = "Tom"
	fmt.Println(reflect.TypeOf(student), student)
	return student
}

// =======================================================================================

type Animal struct {
	Name string
	Run func()
}

type Cat struct {
	Animal
}

func AnimalTest() {
	animal := new(Animal)
	animal.Name = "Dog"
	animal.Run = func() {
		fmt.Println(animal.Name + " is Running...")
	}
	animal.Run()

	cat := new(Cat)
	cat.Name = "Tom"
	cat.Run = func() {
		fmt.Println(cat.Name + "Cat is Running..")
	}
	cat.Run()

}



