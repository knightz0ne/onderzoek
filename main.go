package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseDataFromFile(filename string) ([]map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var records []map[string]string
	current := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		// Separator line = new record
		if strings.HasPrefix(line, "-----") {
			if len(current) > 0 {
				records = append(records, current)
			}
			current = make(map[string]string)
			continue
		}
		// Parse "Key: Value"
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			current[key] = value
		}
	}
	// Add the last block if not empty
	if len(current) > 0 {
		records = append(records, current)
	}
	return records, scanner.Err()
}

func main() {
	filename := "facebook_dataset_small"
	records, err := parseDataFromFile(filename)
	if err != nil {
		panic(err)
	}

	// Extract all passwords
	passwords := []string{}
	for _, rec := range records {
		if pwd, ok := rec["Password"]; ok {
			passwords = append(passwords, pwd)
		}
	}

	// Display the results
	fmt.Println("Found", len(passwords), "passwords:")
	for _, pwd := range passwords {
		fmt.Println(pwd)
	}
}