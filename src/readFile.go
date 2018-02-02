package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No argument, exiting!")
		os.Exit(0)
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()

		scanner := bufio.NewScanner(file)

		counter := 0

		for scanner.Scan() {
			counter++
			fmt.Printf("\nText line %d: %s", counter,scanner.Text())
		}
	}

}

