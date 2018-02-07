package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// To do:
// Exclude filetypes
// Search only certain filetypes
// Eventually, regexes.


type foundElement struct {
	number int
	line   string
}

// Flag used to ask if we print lines or not.
var plns = flag.Bool("printLines", false, "Print the lines where the search term was found")

func main() {
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Word required, exiting!")
		os.Exit(0)
	} else {
		searchterm := flag.Args()[0]
		files := scanDirs()
		for _, f := range files {
			searchFile(f, searchterm)
		}
	}
}

func scanDirs() []string {
	var allDirs []string
	var allFiles []string
	var errors []error
	var currentDir string
	var dirNameToSave string

	allDirs = append(allDirs, ".")

	for len(allDirs) > 0 {
		currentDir = allDirs[0] + "/"
		startFiles, err := ioutil.ReadDir(currentDir)
		// Error checking
		if err != nil {
			errors = append(errors, err)
		} else {
			// Separate between files and dirs, then store in corresponing array
			for _, f := range startFiles {
				if f.IsDir() {
					dirNameToSave = currentDir + f.Name()
					allDirs = append(allDirs, dirNameToSave)
				} else {
					allFiles = append(allFiles, currentDir+f.Name())
				}
			}
		}
		// Slice out the directory we already used
		allDirs = allDirs[1:]
	}

	if len(errors) > 0 {
		fmt.Println("\nErrors:")
		for e := range errors {
			fmt.Println(e)
		}
	}

	return allFiles
}

func searchFile(filename string, search_term string) {
	// initialize variables
	var elementsFound []foundElement
	var elemFound foundElement
	counter := 0
	found := false

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			counter++
			if strings.Contains(scanner.Text(), search_term) {
				found = true
				elemFound.number = counter
				elemFound.line = scanner.Text()
				elementsFound = append(elementsFound, elemFound)
			}
		}

		if found {
			if *plns {
				// print filename, line numbers and lines.
				fmt.Println(filename)
				for _, fElem := range elementsFound {
					fmt.Printf("%d: %s\n", fElem.number, fElem.line)
				}

				fmt.Println("")

			} else {
				// Onl print filename and line numbers
				var lineList string
				for index, fElem := range elementsFound {
					if index == 0 {
						lineList = lineList + strconv.Itoa(fElem.number)
					} else {
						lineList = lineList + ", " + strconv.Itoa(fElem.number)
					}
				}

				fmt.Println(filename + ": " + lineList + "\n")
			}
		}
	}
}
