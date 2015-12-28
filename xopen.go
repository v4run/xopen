package main

import (
	"os"
	"fmt"
	"strings"
	"os/exec"
)

func usage() {
	fmt.Println("\n\033[32;1mXOpen\n\033[33m=====")
	fmt.Println("Usage:\txopen file/directory\n")
}

func main() {
	args := os.Args[1:]
	argLen := len(args)
	if argLen == 0 {
		fmt.Println("\033[31mError: No arguments specified.")
		usage()
		os.Exit(1)
	}
	var file string
	for _, rem := range args {
		file += rem + " "
	}
	file = strings.TrimSpace(file)
	_, fileErr := os.Stat(file)
	if fileErr != nil {
		fmt.Printf("\033[31mError: `%s` not found.\n", file)
		os.Exit(1)
	}
	err := exec.Command("xdg-open", file).Run()
	if err != nil {
		errA := exec.Command("xdg-open", "./" + file).Run()
		if errA != nil {
			fmt.Printf("\033[31mError: %s could not be opened. Make sure you have the right permission.\n", file)
			os.Exit(1)
		}
	}
	fmt.Println("\033[32;1mOpening", file)
}
