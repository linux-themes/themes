package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func install_build() {
	str := strings.Split("mkdir -p bin", " ")
	command := exec.Command(str[0], str[1:]...)
	err := command.Run()
	if err != nil {
		log.Fatalf("command.Run() failed: %v\n", err)
	}

	str = strings.Split("go build -o bin/themes.exe", " ")
	command = exec.Command(str[0], str[1:]...)
	err = command.Run()
	if err != nil {
		log.Fatalf("command.Run() failed: %v\n", err)
	}
	if _, err = os.Stat("./bin/themes.exe"); err != nil {
		log.Fatal(err.Error())
	}
}

func install_local() {
	str := strings.Split("go install", " ")
	command := exec.Command(str[0], str[1:]...)
	err := command.Run()
	if err != nil {
		log.Fatalf("command.Run() failed: %v\n", err)
	}

	home_path, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err.Error())
	}

	if _, err = os.Stat(home_path + "/go/bin/"); err != nil {
		log.Fatal(err.Error())
	}
}

func install_repository() {
	str := strings.Split("go install -x github.com/linux-themes/themes", " ")
	command := exec.Command(str[0], str[1:]...)
	err := command.Run()
	if err != nil {
		log.Fatalf("command.Run() failed: %v\n", err)
	}

	home_path, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err.Error())
	}

	if _, err = os.Stat(home_path + "/go/bin/"); err != nil {
		log.Fatal(err.Error())
	}
}

func Test_Install_Program(t *testing.T) {
	tests := []struct {
		name string
		Test func()
	}{
		{"build", install_build},
		{"install source", install_local},
		{"install repository", install_repository},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
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
		Test func()
	}{
		{"themes list all", list_all},
		{"themes list icons", list_icons},
		{"themes list official", list_themes},
		{"themes list themes", list_official},
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

func Test_Install_Command(t *testing.T) {
	tests := []struct {
		name string
		Test func()
	}{
		{"themes install", install_all},
		{"themes install invalidurl and invalidpackage", install_invalid},
		{"themes install url", install_icons_url},
		{"themes install package", install_themes_url},
		{"themes install url and package", install_icons_official},
		{"themes install invalidurl and invalidpackage", install_themes_official},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			main()
			test.Test()
		})
	}
}

func test_remove_all() {
	str := strings.Split("./bin/themes.exe remove all icons", " ")
	command := exec.Command(str[0], str[1:]...)
	err := command.Run()
	if err != nil {
		log.Fatalf("command.Run() failed: %v\n", err.Error())
	}

	home_path, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Stat(home_path + "/.icons")
	if err == nil {
		log.Fatal(err)
	}

	_, err = os.Stat(home_path + "/.themes")
	if err == nil {
		log.Fatal(err)
	}
}

func test_remove_icons() {
	str := strings.Split("./bin/themes.exe remove all icons", " ")
	command := exec.Command(str[0], str[1:]...)
	err := command.Run()
	if err != nil {
		log.Fatalf("command.Run() failed: %v\n", err.Error())
	}

	home_path, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Stat(home_path + "/.icons")
	if err == nil {
		log.Fatal(err)
	}
}

func test_remove_themes() {
	str := strings.Split("./bin/themes.exe remove all icons", " ")
	command := exec.Command(str[0], str[1:]...)
	err := command.Run()
	if err != nil {
		log.Fatalf("command.Run() failed: %v\n", err.Error())
	}

	home_path, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Stat(home_path + "/.themes")
	if err == nil {
		log.Fatal(err)
	}
}

func Test_Remove_Command(t *testing.T) {
	tests := []struct {
		name string
		Test func()
	}{
		{"themes remove all", test_remove_all},
		{"themes remove icons", test_remove_icons},
		{"themes remove themes", test_remove_themes},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.Test()
		})
	}
}

func Test_Clean(t *testing.T) {
	tests := []struct {
		name string
		Test func()
	}{
		{"clean", list_all},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.Test()
		})
	}
}
