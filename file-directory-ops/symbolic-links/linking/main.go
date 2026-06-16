package main

import (
	"fmt"
	"os"
)

func main() {
	// In the Linux command line, you can create a symbolic link using the ln
	// command with the -s option: ln -s target link
	// ln -s /Users/aythejuggernaut/Downloads/wellat-mobile-logo.svg /Users/aythejuggernaut/Desktop/wellat-logo.svg

	// ln - This is the command for creating links
	// -s - This option is for creating symbolic links
	// target - This is the source file or directory that you want to link to
	// link - This is the destination where you want to create the symbolic link

	// In go, we can use the os.Symlink function to create a symbolic link
	// os.Symlink(target, link)
	// os.Symlink("/Users/aythejuggernaut/Downloads/wellat-mobile-logo.svg", "/Users/aythejuggernaut/Desktop/wellat-logo.svg")
	// Note: The above is not a real file path, just an example

	sourcePath := "/Users/aythejuggernaut/Downloads/wellat-mobile-logo.svg"
	symlinkPath := "/Users/aythejuggernaut/Desktop/wellat-logo.svg"

	// The os.Symlink function is used to create the symlink
	err := os.Symlink(sourcePath, symlinkPath)
	if err != nil {
		fmt.Printf("Error creating symlink: %v\n", err)
		return
	}

	fmt.Printf("Symlink created: %s -> %s\n", symlinkPath, sourcePath)
}
