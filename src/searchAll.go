package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// To do:
// Add help text
// Add regex presets
// Add filename search
// Add ability to search in specified paths
// Add case-insensitive search

type foundElement struct {
	number int
	line   string
}

// Set the search term as a global variable so we can use it accordingly.
var searchTerm string
var compiledRegex *regexp.Regexp

// Flag used to ask if we print lines or not.
var ptlns = flag.Bool("printLines", false, "Print the lines where the search term was found")

// Flag used to ask if we want to search only certain filetypes.
var incExt = flag.String("incExt", "", "Only search in files with the given extension. Multiple inclusions can be given by separating with \",\"")

// Flag to ask if we want to exclude filetypes.
var excExt = flag.String("excExt", "", "Exclude filetypes from search. Multiple exclusions can be given by separating with \",\"")

// Flag to ask if we're going to use a regex.
var regFlag = flag.String("regex", "", "Search based on a regular expression.")

// Help text
var helpFlag = flag.Bool("h", false, "Print this help message")

func main() {
	flag.BoolVar(helpFlag, "help", false, "Print this help message")
	flag.Parse()

	// Make sure either a regex or search term is provided.
	if len(flag.Args()) == 0 && *regFlag == "" {
		files := scanDirs()
		sort.Strings(files)
		fmt.Println("Files found:")
		for _, f := range files {
			fmt.Println(f)
		}
		os.Exit(0)
	} else if len(flag.Args()) > 0 && *regFlag != "" {
		fmt.Println("Please only provide either a search term or regular expression. Exiting.")
		os.Exit(0)
	} else {
		// Assign either search term or compile and assign regex.
		if len(flag.Args()) != 0 {
			searchTerm = flag.Args()[0]
			if len(flag.Args()) > 1 {
				fmt.Println("Remember to put the search term after the flags, otherwise the flags won't work as intended.\n")
			}
		} else {
			fmt.Println(*regFlag)
			var err error
			compiledRegex, err = regexp.Compile(*regFlag)
			if err != nil {
				fmt.Println("Regular expression can't compile, exiting!")
				os.Exit(0)
			}
		}
		// Find files and search them.
		files := scanDirs()
		for _, f := range files {
			searchFile(f)
		}
	}
}

func scanDirs() []string {
	var allDirs []string
	var allFiles []string
	var errors []error
	var currentDir string
	var dirNameToSave string

	// Set the local directory as the starting point to look for more files/directories.
	allDirs = append(allDirs, ".")

	for len(allDirs) > 0 {
		currentDir = allDirs[0] + "/"
		// Returns []os.FileInfo
		startFiles, err := ioutil.ReadDir(currentDir)

		if err != nil {
			errors = append(errors, err)
		} else {
			// Separate between files and dirs
			for _, f := range startFiles {
				if f.IsDir() {
					dirNameToSave = currentDir + f.Name()
					allDirs = append(allDirs, dirNameToSave)
				} else {
					// Checks wether to append the file or not based on filetype arguments
					allFiles = manageExtensions(f.Name(), currentDir, allFiles, *incExt, *excExt)
				}
			}
		}
		// Slice out the directory we already used
		allDirs = allDirs[1:]
	}

	// Print all errors at the end, this is debugging purposes. Probably won't be used.
	if len(errors) > 0 {
		fmt.Println("\n##########################################################")
		fmt.Println("Errors:")
		for e := range errors {
			fmt.Println(e)
		}
		fmt.Println("##########################################################")
	}

	return allFiles
}

// Handles opening the file, passing lines to search and storing hits.
func searchFile(filename string) {
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
			if handleSearch(scanner.Text()) {
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

// Return a list that includes or excludes files using filetype flags. Include overrides exclude.
func manageExtensions(filename string, currentDir string, fileList []string, includingExt string, excludingExt string) []string {
	ext := getFiletype(filename)
	incExtList := strings.Split(includingExt, ",")
	excExtList := strings.Split(excludingExt, ",")

	// If both lists are empty, add.
	if includingExt == "" && excludingExt == "" {
		retArray := append(fileList, currentDir+filename)
		return retArray
		// If extension in include list, add.
	} else if strInArray(ext, incExtList) {
		retArray := append(fileList, currentDir+filename)
		return retArray
		// If extension not in exclude list, add.
	} else if !strInArray(ext, excExtList) && len(excExtList) > 0 {
		retArray := append(fileList, currentDir+filename)
		return retArray
		// return unmodified list since no conditions were met.
	} else {
		return fileList
	}
}

// Split files by "." and return last element in array
func getFiletype(filename string) string {
	fileSplit := strings.Split(filename, ".")
	filetype := fileSplit[len(fileSplit)-1]

	// If the filetype is the same as filename it means that there was no extension.
	if filetype == filename {
		return ""
	} else {
		return filetype
	}
}

// Iterate over an array and check if it contains a string
func strInArray(str string, array []string) bool {
	retBool := false
	for _, s := range array {
		if str == s {
			retBool = true
		}
	}

	return retBool
}

// Perform a search based either on search term or regex.
func handleSearch(scannerText string) bool {
	retVal := false
	if searchTerm != "" {
		if strings.Contains(scannerText, searchTerm) {
			retVal = true
		}
	} else {
		result := compiledRegex.FindString(scannerText)
		if result != "" {
			retVal = true
		}
	}
	return retVal
}
