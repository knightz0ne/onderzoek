package beta

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// -------------------------------------
// Identity Model (same as PlatformFactory)
// -------------------------------------
type Identity struct {
	Program  string
	Url      string
	Login    string
	Password string
	Computer string
	Date     string
	Ip       string
}

// -------------------------------------
// FILE PARSER
// -------------------------------------

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

// -------------------------------------
// CONVERT MAPS -> IDENTITIES
// -------------------------------------

func convertToIdentities(records []map[string]string) []Identity {
	ids := make([]Identity, 0, len(records))

	for _, rec := range records {
		id := Identity{
			Program:  rec["Program"],
			Url:      rec["Url"],
			Login:    rec["Login"],
			Password: rec["Password"],
			Computer: rec["Computer"],
			Date:     rec["Date"],
			Ip:       rec["Ip"],
		}
		ids = append(ids, id)
	}

	return ids
}

// -------------------------------------

func main() {
	filename := "facebook_dataset_small"

	records, err := parseDataFromFile(filename)
	if err != nil {
		panic(err)
	}

	identities := convertToIdentities(records)

	for i, id := range identities {
		fmt.Printf("Identity %d: %+v\n", i+1, id)
	}
}
