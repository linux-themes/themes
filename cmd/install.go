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

func is_valid_url(url string) bool {
	if strings.Contains(url, "https://") &&
		(strings.Contains(url, "tar.xz") || strings.Contains(url, "tar.gz") || strings.Contains(url, "tar.yz")) {
		return true
	}
	return false
}

func download_file(filepath string, url string) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create request: %v\n", err)
		return
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Fprintf(os.Stderr, "received error: %v\n", err)
		return
	}
	defer response.Body.Close()

	f, _ := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	bar := progressbar.NewOptions(
		int(response.ContentLength),
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
	io.Copy(io.MultiWriter(f, bar), response.Body)
}

func extract_tar(filepath string, directory string) error {
	cmd := exec.Command("tar", "-xf", filepath, "-C", directory)
	_, err := cmd.Output()
	if err != nil {
		println(err.Error())
	}
	return err
}

var installIconsCmd = &cobra.Command{
	Use:   "icons",
	Short: "Install icons from community repository or url",
	Long:  `Install icons from community repository or url`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		var valid_packages []string
		var icons = Yaml_get_file("icons")
		var home_path, _ = os.UserHomeDir()

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

		for _, pkg := range args {
			if package_number, err := strconv.Atoi(pkg); err != nil { // if string
				if is_valid_url(pkg) {
					valid_packages = append(valid_packages, pkg)
					println(GREEN + "Valid Package\t\t" + RESET + pkg + RESET)
				}
			} else {
				if package_number > 0 && package_number <= len(icons.Icons) { // if number in range
					valid_packages = append(valid_packages, icons.Icons[pkg].Url)
					println(GREEN + "Valid Package\t\t" + RESET + icons.Icons[pkg].Url + RESET)
				} else {
					println(RED + "Invalid Package\t\t" + RESET + pkg + RESET) // if error
				}
			}
		}

		for _, pkg := range valid_packages {
			file_name := pkg[strings.LastIndex(pkg, "/")+1:]
			println(YELLOW + "Installing\t\t" + RESET + CYAN + pkg + RESET)
			download_file(home_path+"/Downloads/"+file_name, pkg)
			extract_tar(home_path+"/Downloads/"+file_name, home_path+"/.icons")
			println()
		}

		for _, pkg := range valid_packages {
			file_name := pkg[strings.LastIndex(pkg, "/")+1:]
			parts := strings.Split(file_name, ".")
			println(GREEN + "Installed\t\t" + RESET + CYAN + parts[0] + RESET)
		}

	},
}

var installThemesCmd = &cobra.Command{
	Use:   "themes",
	Short: "Install themes from community repository or url",
	Long:  `Install themes from community repository or url`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		var valid_packages []string
		var themes = Yaml_get_file("themes")
		var home_path, _ = os.UserHomeDir()

		if _, err := os.Stat(home_path + "/.themes"); os.IsNotExist(err) {
			if err := os.Mkdir(home_path+"/.themes", os.ModePerm); err != nil {
				log.Fatal(err)
			}
		}

		if _, err := os.Stat(home_path + "/Downloads"); os.IsNotExist(err) {
			if err := os.Mkdir(home_path+"/Downloads", os.ModePerm); err != nil {
				log.Fatal(err)
			}
		}

		for _, pkg := range args {
			if package_number, err := strconv.Atoi(pkg); err != nil { // if string
				if is_valid_url(pkg) {
					valid_packages = append(valid_packages, pkg)
					println(GREEN + "Valid Package\t\t" + RESET + pkg + RESET)
				}
			} else {
				if package_number > 0 && package_number <= len(themes.Themes) { // if number in range
					valid_packages = append(valid_packages, themes.Themes[pkg].Url)
					println(GREEN + "Valid Package\t\t" + RESET + themes.Themes[pkg].Url + RESET)
				} else {
					println(RED + "Invalid Package\t\t" + RESET + pkg + RESET) // if error
				}
			}
		}

		for _, pkg := range valid_packages {
			file_name := pkg[strings.LastIndex(pkg, "/")+1:]
			println(YELLOW + "Installing\t\t" + RESET + CYAN + pkg + RESET)
			download_file(home_path+"/Downloads/"+file_name, pkg)
			extract_tar(home_path+"/Downloads/"+file_name, home_path+"/.themes")
			println()
		}

		for _, pkg := range valid_packages {
			file_name := pkg[strings.LastIndex(pkg, "/")+1:]
			parts := strings.Split(file_name, ".")
			println(GREEN + "Installed\t\t" + RESET + CYAN + parts[0] + RESET)
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
