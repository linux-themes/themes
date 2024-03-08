package commands

import (
	"fmt"
	"os"
	"syscall"
)

func InstallCommand(arguments []string) {
	fmt.Println("\n Installing... ")
	if len(arguments) == 4 {
		if !ValidUrl(arguments[3]) {
			syscall.Exit(1)
		}
		if arguments[2] == ICONS {
			icon_packs := []string{arguments[3]}
			Install(icon_packs, ".icons")
		}
		if arguments[2] == THEMES {
			themes_packs := []string{arguments[3]}
			Install(themes_packs, ".themes")
		}
		if arguments[2] == CONFIG {
			InDevelopment()
			return
		}
		if arguments[2] == "offical" {
			InDevelopment()
			return
		}
		return
	}

	if len(arguments) > 3 {
		urls := arguments[3:]
		directory := arguments[2]

		fmt.Print("Packages: ")
		fmt.Println(urls)

		packages := []string{}
		for _, url := range urls {
			if !ValidUrl(url) {
				HelpCommand()
				fmt.Println("Program End.")
				syscall.Exit(0)
			}
			packages = append(packages, url)
		}
		Install(packages, "."+directory)
		return
	}

	HelpCommand()
}

func Install(links []string, directory string) {
	fmt.Println("\n Creating Directory...")
	err := os.Mkdir(directory, 0777)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, link := range links {
		fmt.Println("\n Installing: " + directory + " ...")

		file_name := StripFileNameGit(link)
		directory_path := BuildPathHomeUserDirectory(directory)
		file_path := directory_path + "/" + file_name

		if err := DownloadFile(file_path, link); err != nil {
			fmt.Println(err.Error())
		}
		if err := Extract_Tar(file_path, directory_path); err != nil {
			fmt.Println(err.Error())
		}
	}
}
