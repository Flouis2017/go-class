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

func FirstTest() {
	// 类似与Java set初始化一个对象
	var student Student
	student.Name = "Flouis"
	student.Age = 27
	student.Gender = "Male"
	student.Hobbies = []string{"Music", "Dota2"}
	//student.Hobbies = append(student.Hobbies, "Music")
	//student.Hobbies = append(student.Hobbies, "Dota2")
	fmt.Println(student)
}

func SecondTest() {
	student := Student{
		Name: "Lily",
		Age: 25,
		Gender: "Female",
		Hobbies: []string{"Singing", "Dancing"},
	}
	fmt.Println(student)
}

func ThirdTest() {
	student := Student{"xxx", 30, "Unknown", []string{"qwer", "asdf"}}
	fmt.Println(student)
}

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



