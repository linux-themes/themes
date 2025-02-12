package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"gopkg.in/yaml.v3"
)

type Data struct {
	Name     string `yaml:"name"`
	Category string `yaml:"category"`
	Url      string `yaml:"url_package"`
}

type List struct {
	Themes  map[string]Data `yaml:"themes"`
	Icons   map[string]Data `yaml:"icons"`
	Configs map[string]Data `yaml:"configs"`
}

func Yaml_print_list(list List) {
	print("themes: hit\n")
	for key, item := range list.Themes {
		print(GREEN+"\nPackage: "+RESET, key)
		print(YELLOW+"\n	Name: "+RESET, item.Name)
		print(YELLOW+"\n	Category: "+RESET, item.Category)
		print(YELLOW+"\ns	Url: "+RESET, item.Url)
		print("\n")
	}
	print("icons: hit\n")
	for key, item := range list.Icons {
		print(GREEN+"\nPackage: "+RESET, key)
		print(YELLOW+"\n	Name: "+RESET, item.Name)
		print(YELLOW+"\n	Category: "+RESET, item.Category)
		print(YELLOW+"\n	Url: "+RESET, item.Url)
		print("\n")
	}
	print("configs: hit\n")
	for key, item := range list.Configs {
		print(GREEN+"\nPackage: "+RESET, key)
		print(YELLOW+"\n	Name: "+RESET, item.Name)
		print(YELLOW+"\n	Category: "+RESET, item.Category)
		print(YELLOW+"\n	Url: "+RESET, item.Url)
		print("\n")
	}
}

func Yaml_print(data Data) {
	print(YELLOW+"\n	Name: "+RESET, data.Name)
	print(YELLOW+"\n	Category: "+RESET, data.Category)
	print(YELLOW+"\n	Url: "+RESET, data.Url)
	print("\n")
}

func Yaml_get_file(folder string) List {
	url := fmt.Sprintf("https://raw.githubusercontent.com/linux-themes/database/refs/heads/main/%s/index.yml", folder)
	resp, err := http.Get(url)
	if err != nil {
		println(url)
		log.Fatalf(RED+"Failed to fetch the file: %v | "+url+RESET, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		println(url)
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

func Yaml_test() {
	data := Yaml_get_file("themes")
	Yaml_print_list(data)
}
