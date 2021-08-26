package array

import (
	"fmt"
)

var Zoo = [...]string{"Tiger", "Lion", "Elephant", "Wolf"}

func Run() {
	for i:=0; i< len(Zoo); i++ {
		fmt.Println(Zoo[i] + " Run...")
	}
}



