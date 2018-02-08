package main

import (
	"flag"
	"io/ioutil"
	"strings"
    "fmt"
)

// Receibe an extension to look for
var ext = flag.String("ext", "", "The extension we're looking for.")

func main() {
    flag.Parse()
    // Get dir files
	dirFiles, err := ioutil.ReadDir(".")
    var fileType string

	if err != nil {
		fmt.Println(err)
	} else {
		for _, file := range dirFiles {
            fileType = getFiletype(file.Name())

            // Compare extension to received extension.
			if fileType == *ext {
                fmt.Println("")
				fmt.Println(file.Name())
			}
		}
	}
}

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