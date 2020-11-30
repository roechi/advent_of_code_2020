package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadLines(filepath string) []string {

	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	s := bufio.NewScanner(file)

	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())

	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}

	return lines
}
