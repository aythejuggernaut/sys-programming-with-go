package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// // Joining file paths
	// dir := "/home/user"
	// file := "document.txt"

	// fullPath := filepath.Join(dir, file)
	// fmt.Println(fullPath)

	// // Cleaning file paths
	// uncleanPath := "/home/user/../documents/file.txt"
	// cleanPath := filepath.Clean(uncleanPath)
	// fmt.Println(cleanPath)

	// Splitting file paths
	path := "/home/user/documents/file.txt"
	dir, file := filepath.Split(path)
	fmt.Println("Directory:", dir)
	fmt.Println("File:", file)

	// Checking if a path matches a pattern
	pattern := "*.txt"
	matched, err := filepath.Match(pattern, file)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Matched:", matched)
}
