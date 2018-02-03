package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Filename and word required, exiting!")
		os.Exit(0)
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()

		search_term := os.Args[2]
		counter := 0
		found := false

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			counter++
			if strings.Contains(scanner.Text(), search_term) {
				found = true
				fmt.Printf("\nLine %d: %s", counter, scanner.Text())
			}
		}

		if !found {
			fmt.Println("Word not found in document.")
		}
	}

}
