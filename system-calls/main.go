package main

import (
	"fmt"
	"io"
	"os"
)

type CliConfig struct {
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
		ErrStream: os.Stderr,
		OutStream: os.Stdout,
	}

	for _, opt := range opts {
		if err := opt(&c); err != nil {
			return CliConfig{}, err
		}
	}

	return c, nil
}

func main() {
	words := os.Args[1:]
	if len(words) == 0 {
		fmt.Fprintf(os.Stderr, "No words provided.\n")
		os.Exit(1)
	}

	cfg, err := NewCliConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating config: %v\n", err)
		os.Exit(1)
	}

	app(words, cfg)
}

func app(words []string, config CliConfig) {
	for _, w := range words {
		if len(w)%2 == 0 {
			fmt.Fprintf(config.OutStream, "word %s is even\n", w)
		} else {
			fmt.Fprintf(config.ErrStream, "word %s is odd\n", w)
		}
	}
}
