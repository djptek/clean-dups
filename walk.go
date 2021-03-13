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
    searchDir := "."

    fileList := []string{}
    //err := 
    if err := filepath.Walk(
        searchDir, 
        func(path string, f os.FileInfo, err error) error {
            fileList = append(fileList, path)
            return nil
        }) ; err != nil {
    }

    for _, file := range fileList {
        fi, err := os.Stat(file)
        if err != nil {
            log.Fatal(err)
        }
        if ! fi.IsDir() {
            data, err := ioutil.ReadFile(file) 
            if err != nil {
                log.Fatal(err)
            }
            sum := sha256.Sum256(data)
            fmt.Printf("%s\t%x\n", file, sum)
        }
    }
}
