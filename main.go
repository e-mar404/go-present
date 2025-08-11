package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Printf("Please only provide one argument, the path to the directory you wish to present\n")
		os.Exit(1)
	}
	
	directoryPath := args[1]

	fmt.Printf("--------------------------\n")
	err := filepath.WalkDir(directoryPath, func(path string, d fs.DirEntry, err error) error {
		// only transform md files to html
		if !d.IsDir() {
			fmt.Printf("file name: %v\n", d.Name())
			fmt.Printf("--------------------------\n")
		}
		return nil
	})

	if err != nil {
		fmt.Printf("There was an error while opening directory on the file path passed: %v\n", err)
	}
}
