package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"gopkg.in/yaml.v3"
)

type Package struct {
	Name    string `yaml:"name"`
	Icon    string `yaml:"icon"`
	Desktop string `yaml:"desktop"`
	Folder  string `yaml:"folder"`
	Path    string `yaml:"path"`
}

type Data struct {
	Name        string  `yaml:"Name"`
	Description string  `yaml:"Description"`
	Category    string  `yaml:"Category"`
	Package     Package `yaml:"Package"`
}

type Index struct {
	Name        string `yaml:"Name"`
	Description string `yaml:"Description"`
	Category    string `yaml:"Category"`
	Icon        string `yaml:"Icon"`
}

type List struct {
	List map[string]Index `yaml:"Themes"`
}

func Yaml_get_file(folder string, file string) Data {
	url := fmt.Sprintf("https://raw.githubusercontent.com/linux-themes/%s/refs/heads/main/%s", folder, file)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf(RED+"Failed to fetch the file: %v"+RESET, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf(RED+"Failed to fetch the file: %s"+RESET, resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf(RED+"Failed to read the response body: %v"+RESET, err)
	}
	data := Data{}
	err = yaml.Unmarshal(body, &data)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return data
}

func Yaml_get_file_index(folder string) List {
	url := fmt.Sprintf("https://raw.githubusercontent.com/linux-themes/%s/refs/heads/main/index.yml", folder)
	resp, err := http.Get(url)
	if err != nil {
		println(url)
		log.Fatalf(RED+"Failed to fetch the file: %v"+RESET, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf(RED+"Failed to fetch the file: %s"+RESET, resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf(RED+"Failed to read the response body: %v"+RESET, err)
	}
	var data List
	err = yaml.Unmarshal([]byte(body), &data)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return data
}

func Yaml_print(data Data) {
	print(GREEN+"\nName: "+RESET, data.Name)
	print(GREEN+"\nDescription: "+RESET, data.Description)
	print(GREEN + "\nPackage:")
	print(YELLOW+"\n    Name: "+RESET, data.Package.Name)
	print(YELLOW+"\n    Icon: "+RESET, data.Package.Icon)
	print(YELLOW+"\n    Desktop: "+RESET, data.Package.Desktop)
	print(YELLOW+"\n    Folder: "+RESET, data.Package.Folder)
	print(YELLOW+"\n    Path: "+RESET, data.Package.Path)
	print("\n")
}

func Yaml_print_list(list List) {
	for key, theme := range list.List {
		print(GREEN+"\nTheme: "+RESET, key)
		print(YELLOW+"\n	Name: "+RESET, theme.Name)
		print(YELLOW+"\n	Description: "+RESET, theme.Description)
		print(YELLOW+"\n	Category: "+RESET, theme.Category)
		print(YELLOW+"\n	Icon: "+RESET, theme.Icon)
		print("\n")
	}
}

func Yaml_print_index(index Index) {
	print(YELLOW+"\n	Name: "+RESET, index.Name)
	print(YELLOW+"\n	Description: "+RESET, index.Description)
	print(YELLOW+"\n	Category: "+RESET, index.Category)
	print(YELLOW+"\n	Icon: "+RESET, index.Icon)
	print("\n")
}

func Yaml_test() {
	index := Yaml_get_file_index(".icons")
	Yaml_print_list(index)

	file := Yaml_get_file(".themes", "data/gnome/marble/marble.yml")
	Yaml_print(file)
}
