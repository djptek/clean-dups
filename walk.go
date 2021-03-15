package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

//func main() ([]string, error) {
func main() {
	searchDir := "test"

	fileList := []string{}
	//err :=
	if err := filepath.Walk(
		searchDir,
		func(path string, f os.FileInfo, err error) error {
			fileList = append(fileList, path)
			return nil
		}); err != nil {
		log.Fatal(err)
	}

	var originals = make(map[string][]string)

	for _, file := range fileList {
		fi, err := os.Stat(file)
		if err != nil {
			log.Fatal(err)
		}
		if !fi.IsDir() {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				log.Fatal(err)
			}
			h := sha256.New()
			h.Write(data)
			var sum = fmt.Sprintf("%s", h.Sum(nil))
			originals[sum] = append(originals[sum], file)
		}
	}

	for sum, files := range originals {
		if len(files) > 1 {
			fmt.Printf("Duplicates:\n")
		} else {
			fmt.Printf("Unique file:\n")
		}
		for _, file := range files {
			fmt.Printf("\t%s\t%x\n", file, sum)
		}
	}

}
