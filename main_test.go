package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func install_program_build() {
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

func install_program_source() {
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

func install_program_repository() {
	str := strings.Split("go install -x github.com/linux-themes/themes@latest", " ")
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
		{"build", install_program_build},
		{"install source", install_program_source},
		// {"install repository", install_program_repository},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.Test()
		})
	}
}

func list_command_all() {
	str := strings.Split("./bin/themes.exe list", " ")
	command := exec.Command(str[0], str[1:]...)
	output, err := command.CombinedOutput()
	if err != nil {
		log.Fatalf("command.Run() failed: %v\n%s\n", err, output)
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

func list_command_icons() {
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

func list_command_themes() {
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

func Test_List_Command(t *testing.T) {
	tests := []struct {
		name string
		Test func()
	}{
		{"themes list", list_command_all},
		{"themes list icons", list_command_icons},
		{"themes list themes", list_command_themes},
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
		"https://github.com/linux-themes/database/raw/refs/heads/main/icons/mint/Mint.tar.xz",
	}
	command := exec.Command(name, str...)
	output, err := command.CombinedOutput()
	if err != nil {
		log.Fatalf("command.Run() failed: %v\n%s\n", err, output)
	}

	home_path, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err.Error())
	}
	if _, err = os.Stat(home_path + "/.icons/Mint"); err != nil {
		log.Fatal(err.Error())
	}
}

func install_themes_url() {
	name := "./bin/themes.exe"
	str := []string{
		"install",
		"themes",
		"https://github.com/linux-themes/database/raw/refs/heads/main/themes/gnome/marble/Marble.tar.gz",
	}
	command := exec.Command(name, str...)
	output, err := command.CombinedOutput()
	if err != nil {
		log.Fatalf("command.Run() failed: %v\n%s\n", err, output)
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

func clean_bin() {
	str := strings.Split("rm -rf bin", " ")
	command := exec.Command(str[0], str[1:]...)
	output, err := command.Output()
	fmt.Printf("%s\n", output)
	if err != nil {
		log.Fatalf("command.Output() failed: %v\n", err.Error())
	}

	_, err = os.Stat("bin")
	if err == nil {
		log.Fatal(err)
	}
}
func clean_go_bin() {
	home_path, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	str := strings.Split("rm "+home_path+"/go/bin/themes", " ")
	command := exec.Command(str[0], str[1:]...)
	output, err := command.Output()
	fmt.Printf("%s\n", output)
	if err != nil {
		log.Fatalf("command.Output() failed: %v\n", err.Error())
	}

	_, err = os.Stat(home_path + "/go/bin/themes")
	if err == nil {
		log.Fatal(err)
	}
}

func Test_Clean(t *testing.T) {
	tests := []struct {
		name string
		Test func()
	}{
		{"clean bin", clean_bin},
		{"clean go bin", clean_go_bin},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.Test()
		})
	}
}
