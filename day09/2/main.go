package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Area struct {
	Position, Length int
}

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
	fragmentedFS, filesAreas := MapToFS(input)
	compactedFS := CompactFS(fragmentedFS, filesAreas)

	//printFS(fragmentedFS)
	//fmt.Println("!!!!!!!!!!!")
	//printFS(compactedFS)

	sum := int64(0)

	for i, v := range compactedFS {
		if v == -1 {
			continue
		}

		sum += int64(i * v)
	}

	return sum
}

func MapToFS(input string) ([]int, map[int]Area) {
	id := 0
	fs := make([]int, 0)
	d := make(map[int]Area)
	for i := 0; i < len(input); {
		n := int(input[i] - '0')

		d[id] = Area{len(fs), n}

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

	return fs, d
}

func CompactFS(fragmentedFS []int, filesAreas map[int]Area) []int {
	compactedFS := make([]int, len(fragmentedFS))
	copy(compactedFS, fragmentedFS)

	key := -1
	for k := range filesAreas {
		if k > key {
			key = k
		}
	}

	for key >= 0 {
		keyLength := filesAreas[key].Length
		keyPosition := filesAreas[key].Position

		startFrom := 0
		for {
			freeArea := FindNextFreeArea(compactedFS, startFrom)

			if freeArea.Position > keyPosition {
				break
			}

			if freeArea.Length >= keyLength {
				for i := range keyLength {
					compactedFS[freeArea.Position+i] = key
					compactedFS[keyPosition+i] = -1
				}

				break
			}

			startFrom = freeArea.Position + freeArea.Length
			if startFrom >= filesAreas[key].Position {
				break
			}
		}

		key--
	}

	return compactedFS
}

func FindNextFreeArea(compactedFS []int, startFrom int) Area {
	startPosition := -1
	length := 0
	for i := startFrom; i < len(compactedFS); i++ {
		if compactedFS[i] == -1 {
			if startPosition == -1 {
				startPosition = i
			}

			length++
		}

		if compactedFS[i] >= 0 && startPosition != -1 {
			break
		}
	}

	area := Area{startPosition, length}
	return area
}

func printFS(fragmentedFS []int) {
	for _, v := range fragmentedFS {
		if v == -1 {
			fmt.Print(".")
		} else {
			fmt.Printf(" %v ", v)
		}
	}
	fmt.Println()
}
