package main

import (
	"fmt"
	//	"strings"
	//	"unicode/utf8"
	//	"unsafe"
)

func main() {

	// Условные операторы
	// if condition {
	//		do'it any
	// } else {
	//		or do'it this
	// }

	var value int
	fmt.Scan(&value)

	if value%2 == 0 {
		fmt.Println("Number ", value, " is even")
	} else {
		fmt.Println("Number ", value, " is odd")
	}

	//END POINT ON VIDEO #2 1:02:20
}
