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
	fmt.Println("Введите число пожалуйста")
	fmt.Scan(&value)

	if value%2 == 0 {
		fmt.Println("Number ", value, " is even")
	} else {
		fmt.Println("Number ", value, " is odd")
	}

	//Инициализация в блоке условного оператора
	//Присваивания и инициализация переменных в условии может быть только через двоеточие и равно, т.е. :=
	//Область видимости инициализируемых переменных ограничена областью видимости условного оператора
	if num := 10; num%2 == 0 {
		fmt.Println("Число: ", num, " ЧЕТНОЕ!")
	} else {
		fmt.Println("Число: ", num, " НЕ ЧЕТНОЕ!")
	}
	//	переменная num - тут уже не будет видна

	//END POINT ON VIDEO #2 1:02:20

}
