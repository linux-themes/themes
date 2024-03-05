package main

import (
	"fmt"
	"os"
)


func main() {
    arguments := os.Args
	if len(arguments) == 0{
		fmt.Println("No arguments error")
		return
	}
	if len(arguments) == 1{
		fmt.Println("Missing arguments error")
		return
	}
	if len(arguments) == 2{
		fmt.Println("Missing arguments error")
		return
	}
	if len(arguments) == 3{
		fmt.Println("Missing arguments error")
		return
	}
	if len(arguments) == 4{
		fmt.Println(arguments[0])
		fmt.Println(arguments[1])
		fmt.Println(arguments[2])
		fmt.Println(arguments[3])
		return
	}
}

