package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("Scanning codebase for BDRA Lite layer boundary leaks...")
	err := filepath.Walk("internal", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.Contains(path, "/pure/") && strings.HasSuffix(path, ".go") {
			file, _ := os.Open(path)
			defer file.Close()
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()
				if strings.Contains(line, "\"database/sql\"") || strings.Contains(line, "\"net/http\"") {
					fmt.Printf("❌ ARCHITECTURE VIOLATION: Banned I/O import leaked into Pure layer: %s\n", path)
					os.Exit(1)
				}
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Linter executed with errors: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✅ Layer boundaries clean. Local lint clearance granted.")
}