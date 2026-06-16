package unlinking

import (
	"fmt"
	"os"
)

func main() {
	// unlink or rm: These commands are used to remove files
	// In a nutshell, link and unlink are the social coordinators of the UNIX
	// filesystem world. link helps make new associations by adding a new name
	// to a file, while unlink sends the file into the oblivion of deletion.
	// They may seem like opposite sides of the same coin, but unlink is the
	// harsh reality check to the merry matchmaking of link.

	filePath := "/Users/aythejuggernaut/Desktop/wellat-logo.svg"

	err := os.Remove(filePath)
	if err != nil {
		fmt.Printf("Error removing the file: %v\n", err)
		return
	}

	fmt.Printf("File removed: %s\n", filePath)
}

// func CheckDanglingSymlink(directories []string, outputWriter *os.File, cfg *Config) {
// 	filePath := "/Users/aythejuggernaut/Desktop/wellat-logo.svg"

// 	for _, directory := range directories {
// 		err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
// 			if err != nil {
// 				fmt.Fprintf(cfg.ErrStream, "Error accessing path %s: %v\n", path, err)
// 				return nil
// 			}
// 			// Check if the current file is a symbolic link.
// 			if info.Mode()&os.ModeSymlink != 0 {
// 				// Resolve the symbolic link.
// 				target, err := os.Readlink(path)
// 				if err != nil {
// 					fmt.Fprintf(cfg.ErrStream, "Error reading symlink %s: %v\n", path, err)
// 				} else {
// 					// Check if the target of the symlink exists.
// 					_, err := os.Stat(target)
// 					if err != nil {
// 						if os.IsNotExist(err) {
// 							fmt.Fprintf(outputWriter, "Broken symlink found: %s -> %s\n", path, target)
// 						} else {
// 							fmt.Fprintf(cfg.ErrStream, "Error checking symlink target %s: %v\n", target, err)
// 						}
// 					}
// 				}
// 			}
// 			return nil
// 		})
// 		if err != nil {
// 			fmt.Fprintf(cfg.ErrStream, "Error walking directory %s: %v\n", directory, err)
// 		}
// 	}
// }

// Let’s break down the most important parts:
// • if info.Mode()&os.ModeSymlink != 0 { ... }: This checks whether the current file is a symbolic link. If it is, it enters this block to resolve and check the validity of the symlink.
// • target, err := os.Readlink(path): This attempts to read the target of the symbolic link using os.Readlink. If an error occurs, it prints an error message indicating that reading the symlink failed.
// • It checks if the target of the symlink exists by using os.Stat(target). If an error occurs during this check, it distinguishes between different types of errors:
// • If the error indicates that the target does not exist (os.IsNotExist(err)), it prints a message indicating a broken symlink.
// • If the error is of another type, it prints an error message indicating that checking the symlink target failed.
