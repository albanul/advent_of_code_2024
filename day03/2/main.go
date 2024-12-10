package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

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

	skip := false

	for _, line := range lines {
		v, s, err := GetMultiplicationSumForLine(line, skip)
		if err != nil {
			return sum, err
		}

		sum += v
		skip = s
	}

	return sum, nil
}

func GetMultiplicationSumForLine(line string, skip bool) (int, bool, error) {
	matches := regexp.MustCompile(`(mul\(\d+,\d+\))|(don't\(\))|(do\(\))`).FindAllString(line, -1)

	sum := 0

	for _, match := range matches {
		if match == "do()" {
			skip = false
			continue
		}

		if match == "don't()" {
			skip = true
			continue
		}

		if skip {
			continue
		}

		v, err := ParseMultiplication(match)
		if err != nil {
			return sum, skip, err
		}

		sum += v
	}

	return sum, skip, nil
}

func ParseMultiplication(s string) (int, error) {
	var x, y int

	_, err := fmt.Sscanf(s, "mul(%d,%d)", &x, &y)
	if err != nil {
		return 0, err
	}

	result := x * y

	return result, nil
}
