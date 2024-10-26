package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"gopkg.in/yaml.v3"
)

type Data struct {
	Name        string `yaml:"Name"`
	Description string `yaml:"Description"`
	Category    string `yaml:"Category"`
	Icon        string `yaml:"Icon"`
}

type List struct {
	List map[string]Data `yaml:"Themes"`
}

func Yaml_get_file(folder string) List {
	url := fmt.Sprintf("https://raw.githubusercontent.com/linux-themes/%s/refs/heads/main/data.yml", folder)
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

func Yaml_print(data Data) {
	print(YELLOW+"\n	Name: "+RESET, data.Name)
	print(YELLOW+"\n	Description: "+RESET, data.Description)
	print(YELLOW+"\n	Category: "+RESET, data.Category)
	print(YELLOW+"\n	Icon: "+RESET, data.Icon)
	print("\n")
}

func Yaml_test() {
	data := Yaml_get_file(".icons")
	Yaml_print_list(data)
}
