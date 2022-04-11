package main

import "fmt"

//for run this programm file, signed string: cd Disc:/path/to/file go run fileName.go
//for exemple this case: cd C:\Users\denis\go\src\github.com\GoLangCouses\Lec2> go run main.go

func main() {
	var a = 9
	var b = 10
	fmt.Println("Result: ", a+b)

	shotCount := 128

	shotCount = shotCount * 3

	fmt.Println("New number variable ShotCount is a: ", shotCount)

}
