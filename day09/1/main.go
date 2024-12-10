package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := GetInputFromFile("day09/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum := CalculateFinalCheckSum(input)
	fmt.Println("Checksum: ", sum)
}

func GetInputFromFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var result string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return result, nil
}

func CalculateFinalCheckSum(input string) int64 {
	fragmentedFS := MapToFS(input)
	compactedFS := CompactFS(fragmentedFS)

	fmt.Println(compactedFS)

	sum := int64(0)

	for i, v := range compactedFS {
		if v == -1 {
			break
		}

		sum += int64(i * v)
	}

	return sum
}

func MapToFS(input string) []int {
	id := 0
	fs := make([]int, 0)
	for i := 0; i < len(input); {
		n := int(input[i] - '0')

		for range n {
			fs = append(fs, id)
		}

		if i+1 == len(input) {
			break
		}

		n = int(input[i+1] - '0')
		for range n {
			fs = append(fs, -1)
		}

		id++
		i += 2
	}

	return fs
}

func CompactFS(fragmentedFS []int) []int {
	l, r := 0, len(fragmentedFS)-1

	compactedFS := make([]int, len(fragmentedFS))
	copy(compactedFS, fragmentedFS)

	for l < r {
		for compactedFS[l] >= 0 {
			l++
		}

		for compactedFS[r] < 0 {
			r--
		}

		if r <= l {
			break
		}

		compactedFS[l], compactedFS[r] = compactedFS[r], compactedFS[l]
	}

	return compactedFS
}
