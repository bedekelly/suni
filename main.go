package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	occurrences := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		occurrences[line]++

		outputBytes, err := json.Marshal(occurrences)
		if err != nil {
			panic(err)
		}

		output := string(outputBytes)

		fmt.Println(output)
	}

	err := scanner.Err()

	if err != nil {
		fmt.Println(err)
	}

}
