package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

const Xmas = "XMAS"

func main() {
	data, err := GetDataFromFile("day04/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	count := CountXmas(data)
	fmt.Println("Count: ", count)
}

func GetDataFromFile(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data := GetData(file)

	return data, nil
}

func GetData(r io.Reader) [][]rune {
	scanner := bufio.NewScanner(r)

	data := make([][]rune, 0)

	for scanner.Scan() {
		temp := make([]rune, 0)

		line := scanner.Text()
		for _, c := range line {
			temp = append(temp, c)
		}

		data = append(data, temp)
	}

	return data
}

func CountXmas(data [][]rune) int {
	count := 0

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if IsXmasUp(data, i, j) {
				count++
			}
			if IsXmasUpRight(data, i, j) {
				count++
			}
			if IsXmasRight(data, i, j) {
				count++
			}
			if IsXmasDownRight(data, i, j) {
				count++
			}
			if IsXmasDown(data, i, j) {
				count++
			}
			if IsXmasDownLeft(data, i, j) {
				count++
			}
			if IsXmasLeft(data, i, j) {
				count++
			}
			if IsXmasUpLeft(data, i, j) {
				count++
			}
		}
	}

	return count
}

func IsXmasUp(data [][]rune, i, j int) bool {
	if i < len(Xmas)-1 {
		return false
	}

	if data[i][j] == 'X' && data[i-1][j] == 'M' && data[i-2][j] == 'A' && data[i-3][j] == 'S' {
		return true
	}

	return false
}

func IsXmasUpRight(data [][]rune, i, j int) bool {
	if i < len(Xmas)-1 || j+len(Xmas) > len(data[0]) {
		return false
	}

	if data[i][j] == 'X' && data[i-1][j+1] == 'M' && data[i-2][j+2] == 'A' && data[i-3][j+3] == 'S' {
		return true
	}

	return false
}

func IsXmasRight(data [][]rune, i, j int) bool {
	if j+len(Xmas) > len(data[0]) {
		return false
	}

	if data[i][j] == 'X' && data[i][j+1] == 'M' && data[i][j+2] == 'A' && data[i][j+3] == 'S' {
		return true
	}

	return false
}

func IsXmasDownRight(data [][]rune, i, j int) bool {
	if i+len(Xmas) > len(data) || j+len(Xmas) > len(data[0]) {
		return false
	}

	if data[i][j] == 'X' && data[i+1][j+1] == 'M' && data[i+2][j+2] == 'A' && data[i+3][j+3] == 'S' {
		return true
	}

	return false
}

func IsXmasDown(data [][]rune, i, j int) bool {
	if i+len(Xmas) > len(data) {
		return false
	}

	if data[i][j] == 'X' && data[i+1][j] == 'M' && data[i+2][j] == 'A' && data[i+3][j] == 'S' {
		return true
	}

	return false
}

func IsXmasDownLeft(data [][]rune, i, j int) bool {
	if i+len(Xmas) > len(data) || j < len(Xmas)-1 {
		return false
	}

	if data[i][j] == 'X' && data[i+1][j-1] == 'M' && data[i+2][j-2] == 'A' && data[i+3][j-3] == 'S' {
		return true
	}

	return false
}

func IsXmasLeft(data [][]rune, i, j int) bool {
	if j < len(Xmas)-1 {
		return false
	}

	if data[i][j] == 'X' && data[i][j-1] == 'M' && data[i][j-2] == 'A' && data[i][j-3] == 'S' {
		return true
	}

	return false
}

func IsXmasUpLeft(data [][]rune, i, j int) bool {
	if i < len(Xmas)-1 || j < len(Xmas)-1 {
		return false
	}

	if data[i][j] == 'X' && data[i-1][j-1] == 'M' && data[i-2][j-2] == 'A' && data[i-3][j-3] == 'S' {
		return true
	}

	return false
}
