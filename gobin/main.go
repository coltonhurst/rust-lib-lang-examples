package main

import (
	"fmt"
	rl "github.com/coltonhurst/rust-lib-lang-examples"
)

func main() {
	result, err := rl.RunCommand(5, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("The sum is:", result)
}
