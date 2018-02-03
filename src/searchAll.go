package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"io/ioutil"
)

// Gotta specify in which file we found the string
// Is using the whole path a good idea? Probably.
// Use a boolean to determine if we print results.

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Word required, exiting!")
		os.Exit(0)
	} else {
		searchterm := os.Args[1]
		files := getFiles()
		for _, f := range files{
			searchFile(f, searchterm)	
		}
	}
}

func getFiles() []string {
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
			// Separate between files and dirs
			for _, f := range startFiles {
				if f.IsDir() {
					dirNameToSave = currentDir + f.Name()
					allDirs = append(allDirs, dirNameToSave)
				} else {
					allFiles = append(allFiles, currentDir + f.Name())
				}
			}
		}
		// Slice out the file we already used
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
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()
		
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