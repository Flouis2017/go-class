package GoArray

import (
	"fmt"
)

var Zoo = [...]string{"Tiger", "Lion", "Elephant", "Wolf", "Bear"}

func Run() {
	for i:=0; i< len(Zoo); i++ {
		fmt.Println(Zoo[i] + " Run...")
	}
}



