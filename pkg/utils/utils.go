package utils

import (
	"errors"
	"fmt"
	"os"

	"github.com/atotto/clipboard"
)

var (
	ErrCopyToClipboard = errors.New("failed to copy to clipboard")
	ErrCreateFile      = errors.New("failed to create file")
	ErrWriteToFile     = errors.New("failed to write to file")
)

// CopyToClipboard copies the given text to the clipboard
func CopyToClipboard(text string) error {
	err := clipboard.WriteAll(text)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrCopyToClipboard, err)
	}
	return nil
}

// SaveToFile saves the given text to a file
func SaveToFile(text, filename string) error {
	// Create file
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrCreateFile, err)
	}

	// Write to file
	_, err = f.WriteString(text)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrWriteToFile, err)
	}

	// Close file
	err = f.Close()
	if err != nil {
		return fmt.Errorf("%w: %v", ErrWriteToFile, err)
	}

	return nil

}
