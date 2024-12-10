package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

const pattern string = `mul\(\d+,\d+\)`

func main() {

	lines, err := GetLinesFromFile("day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum, err := SumMultiplications(lines)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum: ", sum)
}

func GetLinesFromFile(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := GetLines(file)
	return lines, nil
}

func GetLines(reader io.Reader) []string {
	scanner := bufio.NewScanner(reader)

	lines := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func SumMultiplications(lines []string) (int, error) {
	sum := 0

	for _, line := range lines {
		v, err := GetMultiplicationSumForLine(line)
		if err != nil {
			return sum, err
		}

		sum += v
	}

	return sum, nil
}

func GetMultiplicationSumForLine(line string) (int, error) {
	matches := regexp.MustCompile(pattern).FindAllString(line, -1)

	sum := 0

	for _, match := range matches {
		var f, s int

		_, err := fmt.Sscanf(match, "mul(%d,%d)", &f, &s)
		if err != nil {
			return 0, err
		}

		sum += f * s
	}

	return sum, nil
}
