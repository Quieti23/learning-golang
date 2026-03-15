package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	samplePath, err := demoFilePath("sample.txt")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	missingPath, err := demoFilePath("missing.txt")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("success case:")
	if err := printFile(samplePath); err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println()
	fmt.Println("failure case:")
	if err := printFile(missingPath); err != nil {
		fmt.Println("error:", err)
	}
}

func printFile(path string) error {
	name := filepath.Base(path)

	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open %s failed: %w", name, err)
	}
	defer func() {
		fmt.Println("closing file:", name)
		file.Close()
	}()

	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read %s failed: %w", name, err)
	}

	fmt.Printf("content of %s: %s\n", name, content)
	return nil
}

func demoFilePath(name string) (string, error) {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("cannot resolve demo file path")
	}

	return filepath.Join(filepath.Dir(currentFile), name), nil
}
