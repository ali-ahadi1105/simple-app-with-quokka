package main

import (
	"fmt"

	quokka "github.com/ali-ahadi1105/Quokka"
)

func main() {
	result := quokka.TestFunc(2, 3)
	fmt.Println(result)

	result = quokka.TestFunc2(3, 1)
	fmt.Println(result)

	result = quokka.TestFunc3(2, 3)
	fmt.Println(result)
}
