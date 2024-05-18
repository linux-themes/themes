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

const TEST_URL = "https://github.com/sudo-adduser-jordan/mint-y-winx/raw/main/mint-y-winx.tar.xz"

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

func Extract_Tar(filepath string, directory string) error { // add progress bar
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

		for _, packg := range args {
			if offical_package, err := strconv.Atoi(packg); err == nil {
				fmt.Printf("%q looks like a number.\n", packg)
				if offical_package == 1 {
					valid_links = append(valid_links, packages_offical_icons[1].link)
				}

				if offical_package == 2 {
					valid_links = append(valid_links, packages_offical_themes[1].link)
				}
			} else {
				if IsValidUrl(packg) {
					valid_links = append(valid_links, packg)
					fmt.Println(GREEN + "Valid packg:\t" + CYAN + packg + RESET)
				} else {
					fmt.Println(RED + "Invalid packg:\t" + YELLOW + packg + RESET)
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

		is_valid := true
		for _, link := range args { // add community repo source
			if IsValidUrl(link) {
				fmt.Println(GREEN + "Valid link:\t" + CYAN + link + RESET)
			} else {
				is_valid = false
				fmt.Println(RED + "Invalid link:\t" + YELLOW + link + RESET)
			}
		}
		if !is_valid {
			os.Exit(1)
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

		for _, link := range args {
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
	Use:   "install",
	Short: "Install theme",
	Long:  `Install themes from community repository or url`,
	Args:  cobra.ExactArgs(1),
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
