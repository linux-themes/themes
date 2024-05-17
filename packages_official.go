package main

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
