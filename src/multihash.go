// A utility to hash all the files in the current directory with time & date.

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

func main() {

	f, err := os.OpenFile("hash_log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Error abriendo archivo de log. ", err)
	}
	defer f.Close()

	log.SetOutput(f)

	log.Println("Empezando a generar firmas\n")

	all_files := getFiles()
	sort.Strings(all_files)

	for _, filename := range all_files {
		file_hash := hash_256(filename)
		log.Printf("\nGenerando huella para '%s'\nHuella: %s\n\n", filename, file_hash)
	}

	log.Printf("Se finalizó el proceso de generación de %v firmas con éxito\n", len(all_files))
}

func getFiles() []string {
	var allDirs, allFiles []string
	var currentDir string
	var dirNameToSave, fileNameToSave string

	allDirs = append(allDirs, ".")

	for len(allDirs) > 0 {
		currentDir = allDirs[0]
		dirContents, err := ioutil.ReadDir(currentDir)
		// Error checking
		if err != nil {
			log.Fatal(err)
		} else {
			// Separate between files and dirs
			for _, dirElem := range dirContents {
				if dirElem.IsDir() {
					dirNameToSave = currentDir + "/" + dirElem.Name()
					allDirs = append(allDirs, dirNameToSave)
				} else {
					fileNameToSave = currentDir + "/" + dirElem.Name()
					allFiles = append(allFiles, fileNameToSave)
				}
			}
		}
		// Slice out the file we already used
		allDirs = allDirs[1:]
	}

	return allFiles
}

func hash_256(filename string) string {
	file_to_hash, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error abriendo archivo. ", err)
	}
	defer file_to_hash.Close()

	hasher := sha256.New()

	_, err = io.Copy(hasher, file_to_hash)
	if err != nil {
		log.Fatal("Error error copiando archvo al generador de hash. ", err)
	}

	sha := hex.EncodeToString(hasher.Sum(nil))
	return sha
}
