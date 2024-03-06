package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/glamour"
)

func Extract_xz(filepath string) error {
	fmt.Println("Extracting: " + filepath)
	cmd := exec.Command("tar", "-xf", filepath)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(stdout))

	cmd2 := exec.Command("mv", "mint-y-winx", "test/mint-y-winx")
	stdout, err = cmd2.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(stdout))

	if err = os.Remove(filepath); err != nil {
		fmt.Println(err.Error())
	}

	return err
}

func BuildPathUser(directory string) string {
	shell_variables := os.Environ()
	for _, variable := range shell_variables {
		if strings.Contains(variable, "LOGNAME=") {
			current_user := strings.Split(variable, "LOGNAME=")
			return "/home/" + current_user[1] + "/." + directory
		}
	}
	return "build path error"
}

func InDevelopment() {
	file_contents, err := os.ReadFile("markdown/contribute.md")
	if err != nil {
		fmt.Println(err.Error())
	}
	out, err := glamour.Render(string(file_contents), "dark")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(out)
}

func DownloadFile(filepath string, url string) error {
	fmt.Println("Downloading" + filepath)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
