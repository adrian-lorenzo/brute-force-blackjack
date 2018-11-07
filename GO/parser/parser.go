package parser

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

//GetData -
func GetData(path string) []int {
	fd, errorOpen := os.Open(path)
	if errorOpen != nil {
		log.Fatal(errorOpen)
	}

	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	var data []int
	for scanner.Scan() {
		num, errorAtoi := strconv.Atoi(scanner.Text())
		if errorAtoi != nil {
			log.Fatal(errorAtoi)
		}
		data = append(data, num)
	}

	if errorScanner := scanner.Err(); errorScanner != nil {
		log.Fatal(errorScanner)
	}

	return data
}
