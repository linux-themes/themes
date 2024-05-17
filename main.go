package main

import (
	"fmt"
	"io"
	"main/cmd"
	"net/http"
	"os"

	"github.com/schollz/progressbar/v3"
)

// check checks the returned error of a function.
func check(f func() error) {
	if err := f(); err != nil {
		fmt.Fprintf(os.Stderr, "received error: %v\n", err)
	}
}

func main() {
	official_packages := []string{}

	// fmt.Println(GREEN + "Connecting:\t" + CYAN + url + RESET)

	req, _ := http.NewRequest("GET", "github.api", nil)
	resp, _ := http.DefaultClient.Do(req)
	defer check(resp.Body.Close)

	f, _ := os.OpenFile("packages.txt", os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	bar := progressbar.NewOptions(
		int(resp.ContentLength),
		// progressbar.OptionSetDescription(GREEN+"Downloading:\t"+RESET),
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

	fmt.Println(resp.Body)
	fmt.Println(official_packages)

	cmd.Execute()
}
