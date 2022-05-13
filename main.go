package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli/v2"
)

func downloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func handleCli(c *cli.Context) error {

	err := downloadFile(c.String("output"), c.String("url"))

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}

	return nil
}

func main() {
	app := &cli.App{
		Name:  "HTTP Get",
		Usage: "Simple remote file downloader",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "url",
				Aliases: []string{"u"},
				Usage:   "URL to download",
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "path to output file",
			},
		},
		Action: handleCli,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
