package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

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

	for i := 0; i < len(data)-2; i++ {
		for j := 0; j < len(data[i])-2; j++ {
			if IsMM(data, i, j) ||
				IsMS(data, i, j) ||
				IsSM(data, i, j) ||
				IsSS(data, i, j) {
				count++
			}
		}
	}

	return count
}

func IsMM(data [][]rune, i, j int) bool {
	if data[i][j] == 'M' && data[i][j+2] == 'M' &&
		data[i+1][j+1] == 'A' &&
		data[i+2][j] == 'S' && data[i+2][j+2] == 'S' {

		return true
	}

	return false
}

func IsMS(data [][]rune, i, j int) bool {
	if data[i][j] == 'M' && data[i][j+2] == 'S' &&
		data[i+1][j+1] == 'A' &&
		data[i+2][j] == 'M' && data[i+2][j+2] == 'S' {

		return true
	}

	return false
}

func IsSM(data [][]rune, i, j int) bool {
	if data[i][j] == 'S' && data[i][j+2] == 'M' &&
		data[i+1][j+1] == 'A' &&
		data[i+2][j] == 'S' && data[i+2][j+2] == 'M' {

		return true
	}

	return false
}

func IsSS(data [][]rune, i, j int) bool {
	if data[i][j] == 'S' && data[i][j+2] == 'S' &&
		data[i+1][j+1] == 'A' &&
		data[i+2][j] == 'M' && data[i+2][j+2] == 'M' {

		return true
	}

	return false
}
