package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	//Boolean
	var firstBool bool

	fmt.Println(firstBool)

	aBool, bBool := true, false

	fmt.Println("AND: ", aBool && bBool)
	fmt.Println("OR: ", aBool || bBool)
	fmt.Println("NOT: ", !aBool)

	// Numeric Integer
	//Sizes integer type vareables

	var a int = 32
	b := 128

	fmt.Printf("Vareable a type is a: %T and size of: %d bytes\n", a, unsafe.Sizeof(a))
	fmt.Printf("Vareable b type is a: %T and size of: %d bytes\n", b, unsafe.Sizeof(b))

	//Явное приведение типов

	var a32 int32 = 12
	var b64 int64 = 45

	fmt.Println(a32 + int32(b64))
	fmt.Println(int64(a32) + b64)

	fmt.Println(int64(a) + b64)
	// fmt.Println(a + b64) - Без явного преведения типа даже у инта с необъявленной размерностью выдаст ошибку

	// Numeric Float
	//float32 & float64 - У float нужно обязательно указывать разрядность
	floatFirst, floatSecond := 5.78, 47.23

	sum := floatFirst + floatSecond
	sub := floatFirst - floatSecond

	fmt.Printf("Sum: %.3f  AND Sub: %.3f", sum, sub)

	// Numeric Complex
	c1 := complex(5, 7)
	c2 := 12 + 32i
	fmt.Println("Result of complex sum: ")
	fmt.Println(c1 + c2)
	fmt.Println("Result of complex multiply: ")
	fmt.Println(c1 * c2)

	//Тип String - в Go это набор БАЙТ

	nameRu := "Федя"

	nameEn := "Fedya"

	fmt.Println("Length Fedya RUS: ", len(nameRu))
	fmt.Println("Length Fedya ENG: ", len(nameEn))
	fmt.Println("Amount of chars in Fedya RUS: ", utf8.RuneCountInString(nameRu))

	//Rune - Руна это один utfный символ

	concat := nameRu + " " + nameEn

	fmt.Println("Ruesult of concatination: ", concat)

	totalString, subString := "ABCDEFGHJK", "CDE"

	elseSubString := "CBA"

	fmt.Println(strings.Contains(totalString, subString))

	fmt.Println(strings.Contains(totalString, elseSubString))

	//END POINT ON VIDEO #2 0:45:54

	var sampleRune rune
	var anotherRune rune = 'Q' // Для инициализации (присваиванию значения переменной с типом rune)
	// руны символьным значением нужно использовать ковычки ''
	var thirdRune = 234

	fmt.Printf("Rune as empty char: %c\n", sampleRune)
	fmt.Printf("Rune as char at symbol: %c\n", anotherRune)
	fmt.Printf("Rune as char at int32(234): %c\n", thirdRune)

	//Сравнение строк в Go с помощью функции string.Compare()
	fmt.Println("Строки abcd и abcd равны: ", strings.Compare("abcd", "abcd"))
	fmt.Println("Строка abcd > abc: ", strings.Compare("abcd", "abc"))
	fmt.Println("Строка abc < abcd: ", strings.Compare("abc", "abcd"))
	fmt.Println("Строки abcd != ABCD не равны: ", strings.Compare("abcd", "ABCD"))
}
