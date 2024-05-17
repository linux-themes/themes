package main

import (
	"fmt"
	"main/cmd"
)

type Package struct {
	name      string
	link      string
	pack_type string
	selected  bool
}

var packages_offical = map[int]Package{
	1: {"mint-y-winx", "https://", "icons", false},
	2: {"marble-shell", "https://", "themes", false},
}

func main() {
	fmt.Println(packages_offical)
	cmd.Execute()
}
