package main

import (
	"os"
	"path/filepath"
	"fmt"
)

var folder = "fotos_dir"

func main() {
	err := filepath.Walk(folder, walk)
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}

func walk(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}
	
	fmt.Printf("%v \n", path)
	return nil
}
