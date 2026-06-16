package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	directories := []string{".", "/usr/lib"}

	m := map[string]int64{}

	for _, directory := range directories {
		dirSize, err := calculateDirSize(directory)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calculating size of %s: %v\n", directory, err)
			continue
		}
		m[directory] = dirSize
	}

	for dir, size := range m {
		var unit string
		switch {
		case size < 1024:
			unit = "B"
		case size < 1024*1024:
			size /= 1024
			unit = "KB"
		case size < 1024*1024*1024:
			size /= 1024 * 1024
			unit = "MB"
		default:
			size /= 1024 * 1024 * 1024
			unit = "GB"
		}
		fmt.Fprintf(os.Stdout, "%s - %d%s\n", dir, size, unit)
	}
}

func calculateDirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(filePath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !fileInfo.IsDir() {
			size += fileInfo.Size()
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return size, nil
}

func findDuplicateFiles(rootDir string) (map[string][]string, error) {
	duplicates := make(map[string][]string)
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if !info.IsDir() {
			hash, err := computeFileHash(path)
			if err != nil {
				return err
			}
			duplicates[hash] = append(duplicates[hash], path)
		}
		return nil
	})
	return duplicates, err
}

func computeFileHash(filePath string) (string, error) {
	// Attempt to open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	// Ensure that the file is closed when the function exits
	defer file.Close()
	// Create an MD5 hash object
	hash := md5.New()
	// Copy the contents of the file into the hash object
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	// Generate a hexadecimal representation of the MD5 hash and return it
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
