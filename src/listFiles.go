package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Lists all files in the base directory and it's children directories, but doesn't give full path

func main() {
	allF := getFiles()

	fmt.Println("Files:")
	for _, f := range allF {
		fmt.Println(f.Name())
	}
}

func getFiles() []os.FileInfo {
	var allDirs []string
	var allFiles []os.FileInfo
	var errors []error
	var currentDir string
	var dirNameToSave string

	allDirs = append(allDirs, ".")

	for len(allDirs) > 0 {
		currentDir = allDirs[0]
		startFiles, err := ioutil.ReadDir(currentDir)
		// Error checking
		if err != nil {
			errors = append(errors, err)
		} else {
			// Separate between files and dirs
			for _, f := range startFiles {
				if f.IsDir() {
					dirNameToSave = currentDir + "/" + f.Name()
					allDirs = append(allDirs, dirNameToSave)
				} else {
					allFiles = append(allFiles, f)
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
