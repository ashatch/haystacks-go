package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	flags, err := GetNeedleHayFlags()

	if err != nil {
		os.Exit(1)
	}

	needles := make(map[string]bool)

	file, err := os.Open(flags.needles)
	if err != nil {
		fmt.Printf("Could not open %v\n", flags.needles)
		fmt.Printf("Supply --needles /path/to/needles\n\n")
		os.Exit(1)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		needles[fileScanner.Text()] = true
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		_, containsKey := needles[line]

		if containsKey {
			fmt.Println(line)
		}
	}

	if scanErr := scanner.Err(); err != nil {
		log.Fatal(scanErr)
	}
}
