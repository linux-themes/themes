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

func Extract_xz(filepath string, directory string) error {
	fmt.Println("Extracting: " + filepath + " -> " + directory)
	cmd := exec.Command("tar", "-xf", filepath, "-C", directory)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(stdout))

	fmt.Println("Removing: " + filepath)
	if err = os.Remove(filepath); err != nil {
		fmt.Println(err.Error())
	}

	return err
}

func StripFileNameGit(link string) string {
	strings := strings.Split(link, "/")
	return strings[len(strings)-1]
}

func BuildPathHomeUser() string {
	shell_variables := os.Environ()
	for _, variable := range shell_variables {
		if strings.Contains(variable, "LOGNAME=") {
			current_user := strings.Split(variable, "LOGNAME=")
			return "/home/" + current_user[1]
		}
	}
	return "build path error"
}

func BuildPathHomeUserDirectory(directory string) string {
	shell_variables := os.Environ()
	for _, variable := range shell_variables {
		if strings.Contains(variable, "LOGNAME=") {
			current_user := strings.Split(variable, "LOGNAME=")
			return "/home/" + current_user[1] + "/" + directory
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
