package main

import "fmt"

func main() {
	var (
		age  int
		name string
	)

	fmt.Println("Укажите Ваш возраст:")
	fmt.Scan(&age)
	fmt.Println("Введите Ваше имя: ")
	fmt.Scan(&name)

	fmt.Printf("Меня зовут: %s\nМне сейчас %d лет", name, age)

	fmt.Scanln()
}
