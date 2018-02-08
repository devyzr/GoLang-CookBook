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
var ptlns = flag.Bool("printLines", false, "Print the lines where the search term was found")
// Flag used to ask if we want to search only certain filetypes
var incExt = flag.String("incExt", "", "Only search in files with the given extension. Multiple inclusions can be given by separating with \",\"")
// Flag to ask if we want to exclude filetypes
var excExt = flag.String("excExt", "", "Exclude filetypes from search. Multiple exclusions can be given by separating with \",\"")

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
		// Returns []os.FileInfo
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
					allFiles = manageExtensions(f.Name(), currentDir, allFiles, *incExt, *excExt)
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
			if *ptlns {
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

// Return a list that includes files with specified extensions or excludes files with specified extensions. Include overrides exclude.
func manageExtensions(filename string, currentDir string, fileList []string, includingExt string, excludingExt string) []string {
	ext := getFiletype(filename)
	incExtList := strings.Split(includingExt, ",")
	excExtList := strings.Split(excludingExt, ",")

	// If both lists are empty, add
	if includingExt == "" && excludingExt == "" {
		retArray := append(fileList, currentDir + filename)
		return retArray
	// if extension in include list, add
	} else if strInArray(ext, incExtList) {
		retArray := append(fileList, currentDir + filename)
		return retArray
	// if extension not in exclude list, add
	} else if !strInArray(ext, excExtList){
		retArray := append(fileList, currentDir + filename)
		return retArray
	// return unmodified list
	} else {
		return fileList
	}
}

// Split files by "." and return las element in array
func getFiletype (filename string) string {
	fileSplit := strings.Split(filename, ".")
	filetype := fileSplit[len(fileSplit) - 1]

	// If the filetype is the same as filename it means that there was no extension.
	if filetype == filename {
		return ""
	} else {
		return filetype
	}
}

// Iterate over an array and check if it contains a string
func strInArray(str string, array []string) bool{
	retBool := false
	for _, s := range array {
		if str == s {
			retBool = true
		}
	}

	return retBool
}