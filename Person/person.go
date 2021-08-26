package Person

import (
	"fmt"
	"strconv"
)

func GoPrint() {
	fmt.Println("{Name: " + Name + ", Age: " + strconv.Itoa(Age) + ", Gender: " + Gender + ", Company: " + Company + "}")
	fmt.Printf("{name: %s, age: %d, gender: %s, company: %s}\n", Name, Age, Gender, Company)
}

func ToString() string {
	return "{Name=" + Name + ", Age= " + strconv.Itoa(Age) + ", Gender=" + Gender + ", Company=" + Company + "}"
}

func ForPrintNum() {
	fmt.Print("ForPrintNum a: ")
	for a := 0; a < 10; a++ {
		fmt.Printf("%d ", a)
	}

	fmt.Println()
	fmt.Print("ForPrintNum b: ")
	var b = 0
	for b < 10 {
		fmt.Printf("%d ", b)
		b++
	}

	fmt.Println()
	fmt.Print("ForPrintNum c: ")
	for c:=0; c<10; c++ {
		if c == 5 {
			continue
		}
		fmt.Printf("%d ", c)
	}

	fmt.Println()
	fmt.Print("ForPrintNum d: ")
	for d:=0; d<10; d++ {
		if d == 5 {
			break
			//return
		}
		fmt.Printf("%d ", d)
	}
	fmt.Println("\nFor Print Over.")
}
