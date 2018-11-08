package parser

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
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

func Shuffle(data []int) []int {
	var result []int
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))
	indexes := randomizer.Perm(len(data))
	for _, index := range indexes {
		result = append(result, data[index])
	}
	return result
}
