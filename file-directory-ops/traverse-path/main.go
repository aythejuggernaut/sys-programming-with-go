package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type CliConfig struct {
	OutputFile           string
	ErrStream, OutStream io.Writer
}

type Option func(config *CliConfig) error

func WithErrorStream(errStream io.Writer) Option {
	return func(config *CliConfig) error {
		if errStream == nil {
			return fmt.Errorf("error stream is nil")
		}
		config.ErrStream = errStream
		return nil
	}
}

func WithOutStream(outStream io.Writer) Option {
	return func(config *CliConfig) error {
		if outStream == nil {
			return fmt.Errorf("output stream is nil")
		}
		config.OutStream = outStream
		return nil
	}
}

func NewCliConfig(opts ...Option) (CliConfig, error) {
	c := CliConfig{
		OutputFile: "",
		ErrStream:  os.Stderr,
		OutStream:  os.Stdout,
	}

	for _, opt := range opts {
		if err := opt(&c); err != nil {
			return CliConfig{}, err
		}
	}

	return c, nil
}

func main() {
	var outputFileName string
	flag.StringVar(&outputFileName, "f", "", "Output file (default: stdout)")
	flag.Parse()

	directories := os.Args[1:]
	if len(directories) == 0 {
		fmt.Fprintf(os.Stderr, "No directories provided.\n")
		os.Exit(1)
	}

	cfg, err := NewCliConfig(WithOutStream(os.Stdout), WithErrorStream(os.Stderr))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating config: %v\n", err)
		os.Exit(1)
	}

	app(directories, cfg)
}

func app(directories []string, config CliConfig) {
	var outputWriter io.Writer
	if config.OutputFile != "" {
		outputFile, err := os.Create(config.OutputFile)
		if err != nil {
			fmt.Fprintf(config.ErrStream, "Error creating output file: %v\n", err)
			os.Exit(1)
		}
		defer outputFile.Close()
		outputWriter = io.MultiWriter(config.OutStream, outputFile)
	} else {
		outputWriter = config.OutStream
	}

	for _, directory := range directories {
		err := filepath.WalkDir(directory, func(path string, d os.DirEntry, err error) error {
			if path == ".git" {
				return filepath.SkipDir
			}
			if d.IsDir() {
				fmt.Fprintf(outputWriter, "%s\n", path)
			}
			return nil
		})
		if err != nil {
			fmt.Fprintf(config.ErrStream, "Error walking the path %q: %v\n", directory, err)
			continue
		}
	}
}

// func main() {
// 	rootPath := "."

// 	err := filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
// 		if err != nil {
// 			fmt.Println("Error accessing path:", path, err)
// 			return err
// 		}

// 		fmt.Println("Path:", path)
// 		return nil
// 	})

// 	if err != nil {
// 		fmt.Println("Error walking the path:", err)
// 		os.Exit(1)
// 	}
// }
