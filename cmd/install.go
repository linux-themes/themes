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

var offical_icons = map[int]string{
	1: "https://github.com/linux-themes/themes-official/raw/main/icons/Infinity.tar.gz",
	2: "https://github.com/linux-themes/themes-official/raw/main/icons/Obsidian.tar.xz",
	3: "https://github.com/linux-themes/themes-official/raw/main/icons/We10X.tar.xz",
	4: "https://github.com/linux-themes/themes-official/raw/main/icons/WhiteSur.tar.xz",
	5: "https://github.com/linux-themes/themes-official/raw/main/icons/Win10Sur.tar.xz",
	6: "https://github.com/linux-themes/themes-official/raw/main/icons/Win11.tar.xz",
	7: "https://github.com/linux-themes/themes-official/raw/main/icons/Mint.tar.xz",
}
var offical_themes = map[int]string{
	1: "https://github.com/linux-themes/themes-official/raw/main/themes/gnome/MarbleShell.tar.gz",
}

func IsValidUrl(url string) bool {
	if strings.Contains(url, "https://") &&
		(strings.Contains(url, "tar.xz") || strings.Contains(url, "tar.gz") || strings.Contains(url, "tar.yz")) {
		return true
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

	fmt.Println(GREEN + "Connecting:\t" + CYAN + url + RESET)

	req, _ := http.NewRequest("GET", url, nil)
	resp, _ := http.DefaultClient.Do(req)
	defer check(resp.Body.Close)

	f, _ := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	bar := progressbar.NewOptions(
		int(resp.ContentLength),
		progressbar.OptionSetDescription(GREEN+"Downloading:\t"+RESET),
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

		valid_links := []string{}
		for _, link := range args {
			if offical_package, err := strconv.Atoi(link); err == nil {
				if offical_package <= len(offical_icons) {
					valid_package := offical_icons[offical_package]
					valid_links = append(valid_links, valid_package)
					fmt.Println(GREEN + "Valid Package:\t" + CYAN + valid_package + RESET)
				} else {
					fmt.Println(RED + "Invalid Package:\t" + YELLOW + link + RESET)
				}
			} else {
				if IsValidUrl(link) {
					valid_links = append(valid_links, link)
					fmt.Println(GREEN + "Valid link:\t" + CYAN + link + RESET)
				} else {
					fmt.Println(RED + "Invalid link:\t" + YELLOW + link + RESET)
				}
			}
		}

		home_path, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		download_path := home_path + "/Downloads"
		install_path := home_path + "/.icons"

		if _, err := os.Stat(install_path); os.IsNotExist(err) {
			if err := os.Mkdir(install_path, os.ModePerm); err != nil {
				log.Fatal(err)
			}
		}

		if _, err := os.Stat(download_path); os.IsNotExist(err) {
			if err := os.Mkdir(download_path, os.ModePerm); err != nil {
				log.Fatal(err)
			}
		}

		for _, link := range valid_links {
			fmt.Println()
			last_index := strings.LastIndex(link, "/")
			file_name := link[last_index+1:]
			fmt.Println(GREEN + "Installing:\t" + RESET + CYAN + install_path + "/" + file_name + RESET)
			DownloadFile(download_path+"/"+file_name, link)
			Extract_Tar(download_path+"/"+file_name, install_path)
			parts := strings.Split(file_name, ".")
			fmt.Println(GREEN + "Installed:\t" + RESET + CYAN + parts[0] + RESET)
		}
	},
}

var installThemesCmd = &cobra.Command{
	Use:   "themes",
	Short: "Install themes from community repository or url",
	Long:  `Install themes from community repository or url`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		valid_links := []string{}
		for _, link := range args {
			if offical_package, err := strconv.Atoi(link); err == nil {
				if offical_package <= len(offical_themes) {
					valid_package := offical_themes[offical_package]
					valid_links = append(valid_links, valid_package)
					fmt.Println(GREEN + "Valid Package:\t" + CYAN + valid_package + RESET)
				} else {
					fmt.Println(RED + "Invalid Package:\t" + YELLOW + link + RESET)
				}
			} else {
				if IsValidUrl(link) {
					valid_links = append(valid_links, link)
					fmt.Println(GREEN + "Valid link:\t" + CYAN + link + RESET)
				} else {
					fmt.Println(RED + "Invalid link:\t" + YELLOW + link + RESET)
				}
			}
		}

		home_path, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		download_path := home_path + "/Downloads"
		install_path := home_path + "/.themes"

		if _, err := os.Stat(install_path); os.IsNotExist(err) {
			if err := os.Mkdir(install_path, os.ModePerm); err != nil {
				log.Fatal(err)
			}
		}

		for _, link := range valid_links {
			last_index := strings.LastIndex(link, "/")
			file_name := link[last_index+1:]
			fmt.Println(GREEN + "Installing:\t" + RESET + CYAN + install_path + "/" + file_name + RESET)
			DownloadFile(download_path+"/"+file_name, link)
			Extract_Tar(download_path+"/"+file_name, install_path)
			parts := strings.Split(file_name, ".")
			fmt.Println(GREEN + "Installed:\t" + RESET + CYAN + parts[0] + RESET)
		}
	},
}

var installCmd = &cobra.Command{
	Use:       "install",
	Short:     "Install theme",
	Long:      `Install themes from community repository or url`,
	ValidArgs: []string{"icons", "themes"},
	Args:      cobra.OnlyValidArgs,
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
