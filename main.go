package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	showIntro()

	for {
		showMenu()
		command := readCommand()

		switch command {
		case 1:
			initMonitoring()
		case 2:
			fmt.Println("Showing Logs...")
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Unknown command")
			os.Exit(-1)
		}
	}
}

func showIntro() {
	const name = "JP"
	const version = 1.1
	fmt.Println("Hi,", name)
	fmt.Println("Version", version)
}

func showMenu() {
	fmt.Println("1. Start Monitoring")
	fmt.Println("2. Show Logs")
	fmt.Println("0. Exit")
}

func readCommand() int {
	var command int

	fmt.Println("Command:")
	fmt.Scan(&command)

	return command
}

func initMonitoring() {
	fmt.Println("Monitoring...")

	sites := readSitesFromFile()

	for _, site := range sites {
		testSite(site)
	}
}

func testSite(site string) {
	response, _ := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println(site, "OK")
	} else {
		fmt.Println(site, "Failed", response.StatusCode)
	}
}

func readSitesFromFile() []string {
	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Error:", err)
		return sites
	}

	reader := bufio.NewReader(file)

	for {
		row, err := reader.ReadString('\n')
		row = strings.TrimSpace(row)
		sites = append(sites, row)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}
