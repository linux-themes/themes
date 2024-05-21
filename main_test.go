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

func list_all() {
	str := strings.Split("./bin/themes.exe list", " ")
	command := exec.Command(str[0], str[1:]...)
	err := command.Run()
	if err != nil {
		log.Fatalf("command.Run() failed: %v\n", err)
	}

	home_path, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err.Error())
	}

	if _, err = os.Stat(home_path + "/.icons"); err != nil {
		log.Fatal(err.Error())
	}

	if _, err = os.Stat(home_path + "/.themes"); err != nil {
		log.Fatal(err.Error())
	}
}
func list_icons() {
	str := strings.Split("./bin/themes.exe list icons", " ")
	command := exec.Command(str[0], str[1:]...)
	err := command.Run()
	if err != nil {
		log.Fatalf("command.Run() failed: %v\n", err)
	}

	home_path, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err.Error())
	}

	if _, err = os.Stat(home_path + "/.icons"); err != nil {
		log.Fatal(command)
		log.Fatal(err.Error())
	}
}

func list_themes() {
	str := strings.Split("./bin/themes.exe list themes", " ")
	command := exec.Command(str[0], str[1:]...)
	err := command.Run()
	if err != nil {
		log.Fatalf("command.Run() failed: %v\n", err)
	}

	home_path, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err.Error())
	}

	if _, err = os.Stat(home_path + "/.themes"); err != nil {
		log.Fatal(err.Error())
	}
}

func Test_List(t *testing.T) {
	tests := []struct {
		name string
		Test func()
	}{
		{"themes list", list_all},
		{"themes list icons", list_icons},
		{"themes list themes", list_themes},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.Test()
		})
	}
}

func install_icons_url() {
	name := "./bin/themes.exe"
	str := []string{
		"install",
		"icons",
		"https://github.com/linux-themes/themes-official/raw/main/icons/mint.tar.xz",
	}
	command := exec.Command(name, str...)
	err := command.Run()
	if err != nil {
		log.Fatalf("command.Run() failed: %v\n", err)
	}

	home_path, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err.Error())
	}
	if _, err = os.Stat(home_path + "/.icons/mint"); err != nil { //fix
		log.Fatal(err.Error())
	}
}

func install_themes_url() {
	name := "./bin/themes.exe"
	str := []string{
		"install",
		"themes",
		"https://github.com/linux-themes/themes-official/raw/main/themes/MarbleShell.tar.gz",
	}
	command := exec.Command(name, str...)
	err := command.Run()
	if err != nil {
		log.Fatalf("command.Run() failed: %v\n", err)
	}

	home_path, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err.Error())
	}
	if _, err = os.Stat(home_path + "/.themes/Marble-blue-dark"); err != nil {
		log.Fatal(err.Error())
	}
}

func Test_Install_Command(t *testing.T) {
	tests := []struct {
		name string
		Test func()
	}{
		{"themes install icons url", install_icons_url},
		{"themes install themes url", install_themes_url},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.Test()
		})
	}
}

func test_remove_all() {
	str := strings.Split("./bin/themes.exe remove all", " ")
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
	str := strings.Split("./bin/themes.exe remove all themes", " ")
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
