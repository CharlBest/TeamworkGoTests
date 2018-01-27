package customerimporter

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
	Importer 2:
	This version was created because of the chance that reader.ReadAll() could use lots 
	of memory and by only reading single etries it would be more efficient
*/
func Importer2() {
	// Load file.
	file, err := os.Open("./customers.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Create a new reader.
	reader := csv.NewReader(bufio.NewReader(file))
	emails := make(map[string]int)

	for {
		//Prevent reader.ReadAll() because of memory implications
		line, err := reader.Read()
		// Stop at end-of-file
		if err == io.EOF {
			break
		}

		components := strings.Split(line[2], "@")

		//Check if entry was split in 2 (correct email address)
		if len(components) > 1 {
			emails[components[1]]++
		}
	}

	fmt.Println(emails)
}
