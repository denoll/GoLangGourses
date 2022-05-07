package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	logo := ""
	var countChars int
	fmt.Scan(&logo)
	countChars = utf8.RuneCountInString(logo)
	price := 23 * countChars
	if price < 100 {
		fmt.Println(price, " коп.")
	} else {
		rub := price / 100
		kop := price % 100
		fmt.Println(rub, " р. ", kop, " коп.")
	}
}
