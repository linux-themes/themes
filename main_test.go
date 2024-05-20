package main

import (
	"log"
	"os"
	"testing"
)

type TestFunction func()

func test_remove_all() {
	home_path, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Stat(home_path + "/.icons")
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Stat(home_path + "/.themes")
	if err != nil {
		log.Fatal(err)
	}
}

func test_remove_icons() {
	home_path, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Stat(home_path + "/.icons")
	if err != nil {
		log.Fatal(err)
	}
}

func test_remove_themes() {
	home_path, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Stat(home_path + "/.themes")
	if err != nil {
		log.Fatal(err)
	}
}

func Test_Remove(t *testing.T) {
	tests := []struct {
		name string
		Test TestFunction
	}{
		{"Test: themes remove all", test_remove_all},
		{"Test: themes remove icons", test_remove_icons},
		{"Test: themes remove themes", test_remove_themes},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			main()
			test.Test()
		})
	}
}

func list_all()      {}
func list_icons()    {}
func list_themes()   {}
func list_official() {}

func Test_List(t *testing.T) {
	tests := []struct {
		name string
		Test TestFunction
	}{
		{"Test: themes list all", list_all},
		{"Test: themes list icons", list_icons},
		{"Test: themes list official", list_themes},
		{"Test: themes list themes", list_official},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			main()
			test.Test()
		})
	}
}

func set_all()    {}
func set_icons()  {}
func set_themes() {}

func Test_Set(t *testing.T) {
	tests := []struct {
		name string
		Test TestFunction
	}{
		{"Test: themes set", set_all},
		{"Test: themes set icons", set_icons},
		{"Test: themes set themes", set_themes},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			main()
			test.Test()
		})
	}
}

func install_all()             {}
func install_icons_url()       {}
func install_themes_url()      {}
func install_icons_official()  {}
func install_themes_official() {}
func install_invalid()         {}

func Test_Install(t *testing.T) {
	tests := []struct {
		name string
		Test TestFunction
	}{
		{"Test: themes install", install_all},
		{"Test: themes install invalidurl and invalidpackage", install_invalid},
		{"Test: themes install url", install_icons_url},
		{"Test: themes install package", install_themes_url},
		{"Test: themes install url and package", install_icons_official},
		{"Test: themes install invalidurl and invalidpackage", install_themes_official},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			main()
			test.Test()
		})
	}
}
