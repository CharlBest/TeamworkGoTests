package customerimporter

import (
	"encoding/csv"
	"fmt"
	"os"
	//"regexp"
	"strings"
)

func extractEmailDomain(email string) (string, error) {
	components := strings.Split(email, "@")
	//I find regex to be slower than Split
	//components := regex.FindStringSubmatch(line[2])

	// Prevent errors if email is not valid (needs @ sign) (3 entries are "email")
	if len(components) > 1 {
		return components[1], nil
	}

	//return "", errors.New("Email adrress is not in valid format")
	return "", nil
}

func parseLines(lines [][]string) map[string]int {
	emails := make(map[string]int)
	//regex := regexp.MustCompile("^.+@(.+$)")

	for _, line := range lines {
		domain, err := extractEmailDomain(line[2])
		if err == nil {
			emails[domain]++
		}
	}

	return emails
}

func Importer() {
	// Loading csv file
	file, err := os.Open("./customers.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Creating csv reader
	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(parseLines(lines))
}
