package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

func IsValidUrl(url string) bool {
	if strings.Contains(url, "https://") &&
		(strings.Contains(url, "tar.xz") || strings.Contains(url, "tar.gz") || strings.Contains(url, "tar.yz")) {
		return true
	}
	return false
}

func IsValidStorePackage(list List, link string) bool {
	num, err := strconv.Atoi(link)
	if err == nil { // If number
		if num < 1 || num > 40 {
			fmt.Printf("number %d is out of range (%d - %d)", num, 1, 40)
			return false
		}
		return true
	}

	for _, data := range list.List { // If package name matches
		if strings.ToLower(data.Name) == link {
			return true
		}
	}
	return false
}

// check checks the returned error of a function.
func check(f func() error) {
	if err := f(); err != nil {
		fmt.Fprintf(os.Stderr, "received error: %v\n", err)
	}
}

func DownloadFile(filepath string, url string) {

	fmt.Println(GREEN + "Connecting\t" + CYAN + url + RESET)

	req, _ := http.NewRequest("GET", url, nil)
	resp, _ := http.DefaultClient.Do(req)
	defer check(resp.Body.Close)

	f, _ := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	bar := progressbar.NewOptions(
		int(resp.ContentLength),
		progressbar.OptionSetDescription(GREEN+"Downloading\t"+RESET),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionUseANSICodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetTheme(
			progressbar.Theme{
				Saucer: "[green]|[reset]",
			},
		),
	)
	io.Copy(io.MultiWriter(f, bar), resp.Body)

	fmt.Println()
}

func Extract_Tar(filepath string, directory string) error {
	cmd := exec.Command("tar", "-xf", filepath, "-C", directory)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

var installIconsCmd = &cobra.Command{
	Use:   "icons",
	Short: "Install icons from community repository or url",
	Long:  `Install icons from community repository or url`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		icons := Yaml_get_file("icons")

		valid_links := []string{}
		for _, link := range args {
			if IsValidStorePackage(icons, link) {
				if _, err := strconv.Atoi(link); err == nil {
					valid_links = append(valid_links, icons.List[link].Url)
					fmt.Println(GREEN + "Valid Package\t\t" + RESET + icons.List[link].Url + RESET)
				} else {
					for _, theme := range icons.List {
						if strings.ToLower(theme.Name) == link {
							fmt.Println(GREEN + "Valid Package\t\t" + RESET + theme.Url + RESET)
							valid_links = append(valid_links, theme.Url)
						}
					}
				}
			} else if IsValidUrl(link) {
				valid_links = append(valid_links, link)
				fmt.Println(GREEN + "Valid Package\t" + RESET + link + RESET)
			} else {
				fmt.Println(RED + "Invalid Package\t" + RESET + link + RESET)
			}
		}

		home_path, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}

		if _, err := os.Stat(home_path + "/.icons"); os.IsNotExist(err) {
			if err := os.Mkdir(home_path+"/.icons", os.ModePerm); err != nil {
				log.Fatal(err)
			}
		}

		if _, err := os.Stat(home_path + "/Downloads"); os.IsNotExist(err) {
			if err := os.Mkdir(home_path+"/Downloads", os.ModePerm); err != nil {
				log.Fatal(err)
			}
		}

		for _, link := range valid_links {
			last_index := strings.LastIndex(link, "/")
			file_name := link[last_index+1:]
			fmt.Println(GREEN + "Installing\t" + RESET + CYAN + home_path + "/.icons/" + file_name + RESET)
			DownloadFile(home_path+"/Downloads/"+file_name, link)
			Extract_Tar(home_path+"/Downloads/"+file_name, home_path+"/.icons")
			parts := strings.Split(file_name, ".")
			fmt.Println(GREEN + "Installed\t" + RESET + CYAN + parts[0] + RESET)
		}
	},
}

var installThemesCmd = &cobra.Command{
	Use:   "themes",
	Short: "Install themes from community repository or url",
	Long:  `Install themes from community repository or url`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		themes := Yaml_get_file("themes")

		valid_links := []string{}
		for _, link := range args {
			if IsValidStorePackage(themes, link) {
				if _, err := strconv.Atoi(link); err == nil {
					valid_links = append(valid_links, themes.List[link].Url)
					fmt.Println(GREEN + "Valid Package\t\t" + RESET + themes.List[link].Url + RESET)
				} else {
					for _, theme := range themes.List {
						if strings.ToLower(theme.Name) == link {
							fmt.Println(GREEN + "Valid Package\t\t" + RESET + theme.Url + RESET)
							valid_links = append(valid_links, theme.Url)
						}
					}
				}
			} else if IsValidUrl(link) {
				valid_links = append(valid_links, link)
				fmt.Println(GREEN + "Valid Package\t" + RESET + link + RESET)
			} else {
				fmt.Println(RED + "Invalid Package\t" + RESET + link + RESET)
			}
		}

		home_path, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}

		if _, err := os.Stat(home_path + "/.themes"); os.IsNotExist(err) {
			if err := os.Mkdir(home_path+"/.themes", os.ModePerm); err != nil {
				log.Fatal(err)
			}
		}

		for _, link := range valid_links {
			last_index := strings.LastIndex(link, "/")
			file_name := link[last_index+1:]
			fmt.Println(GREEN + "Installing\t\t" + RESET + CYAN + home_path + "/.themes/" + file_name + RESET)
			DownloadFile(home_path+"/Downloads/"+file_name, link)
			Extract_Tar(home_path+"/Downloads/"+file_name, home_path+"/.themes/")
			parts := strings.Split(file_name, ".")
			fmt.Println(GREEN + "Installed\t\t" + RESET + CYAN + parts[0] + RESET)
		}
	},
}

var installCmd = &cobra.Command{
	Use:       "install",
	Short:     "Install theme",
	Long:      `Install themes from community repository or url`,
	ValidArgs: []string{"icons", "themes"},
	Args:      cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		if args[0] == "icons" {
			installIconsCmd.Run(cmd, args)
		}

		if args[0] == "themes" {
			installThemesCmd.Run(cmd, args)
		}

	},
}

func init() {
	installCmd.AddCommand(installIconsCmd)
	installCmd.AddCommand(installThemesCmd)
	rootCmd.AddCommand(installCmd)
}
