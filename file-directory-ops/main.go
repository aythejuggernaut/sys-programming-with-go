package main

import (
	"fmt"
	"os"
)

func main() {
	info, err := os.Stat("example.txt")

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
		} else if os.IsPermission(err) {
			fmt.Println("Permission denied")
		} else {
			panic(err)
		}
	}

	permissions := info.Mode().Perm()
	permissionString := fmt.Sprintf("%o", permissions)

	fmt.Printf("File name: %s\n", info.Name())
	fmt.Printf("File size: %d bytes\n", info.Size())
	fmt.Printf("Is directory: %v\n", info.IsDir())
	fmt.Printf("Last modified: %v\n", info.ModTime())
	fmt.Printf("Permissions: %v\n", info.Mode().Perm().String())
	fmt.Printf("Octal permissions: %v\n", permissionString)
	fmt.Printf("System mod: %v\n", info.Sys())
}
