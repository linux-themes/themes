package main

import (
	"fmt"
	"os"
)

const HELP = "help"
const INIT = "init"
const LIST = "list"
const INSTALL = "install"
const REMOVE = "remove"

func main() {
	arguments := os.Args
	switch len(arguments) {
	case 0:
		fmt.Println("No Arguments Error")
	case 4:
		executeArguments(arguments)
	default:
		fmt.Println("Invalid Arguments Error")
	}
}

func executeArguments(arguments []string) {
	fmt.Print("Execute: ")
	switch arguments[1] {
	case HELP:
		fmt.Println(arguments[1])
	case INIT:
		fmt.Println(arguments[1])
	case LIST:
		fmt.Println(arguments[1])
	case INSTALL:
		fmt.Println(arguments[1])
		install(arguments[1], 0)
	case REMOVE:
		fmt.Println(arguments[1])
	default:
		fmt.Println("Execute Arguments Error")
	}
}

func install(packg string, packg_id int) {
	fmt.Println(packg)
	fmt.Println(packg_id)
}
