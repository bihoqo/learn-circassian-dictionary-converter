package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

// ReadFileLineByLine reads a file and executes a callback for each line.
// If the callback returns an error, processing stops and that error is returned.
func ReadFileLineByLine(filePath string, callback func(line string, index int) error) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	// Ensure file is closed even if an error occurs
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("close file error: %v", err)
		}
	}(f)

	sc := bufio.NewScanner(f)

	// Increase buffer to 1MB to handle potential long dictionary entries
	const maxCapacity = 1024 * 1024
	buf := make([]byte, 0, 64*1024)
	sc.Buffer(buf, maxCapacity)

	idx := 0
	for sc.Scan() {
		if err := callback(sc.Text(), idx); err != nil {
			return err // Stop and propagate error from the callback
		}
		idx++
	}

	return sc.Err() // Return scanner errors (like file read issues)
}

// SaveDictToJSON marshals the object to JSON, applies specific string replacements
// (un-escaping HTML and swapping \" for '), and writes to the specific file path.
func SaveDictToJSON(filePath string, data interface{}) error {
	// 1. Marshal with indentation
	bytes, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return fmt.Errorf("json marshal error: %w", err)
	}

	// 2. Convert to string to perform replacements
	jsonStr := string(bytes)

	// 3. Apply the specific replacements from your old repo logic
	// specific to keeping HTML readable and cleaning up quotes
	jsonStr = strings.ReplaceAll(jsonStr, "\\u003c", "<")
	jsonStr = strings.ReplaceAll(jsonStr, "\\u003e", ">")
	jsonStr = strings.ReplaceAll(jsonStr, "\\u0026", "&") // Good practice to unescape ampersands too
	jsonStr = strings.ReplaceAll(jsonStr, "\\\"", "'")

	// 4. Create and Write to file
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("create file error: %w", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("close file error: %v", err)
		}
	}(f)

	_, err = f.WriteString(jsonStr)
	if err != nil {
		return fmt.Errorf("write file error: %w", err)
	}

	return nil
}
