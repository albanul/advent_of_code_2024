package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type InputValues struct {
	LeftList  []int
	RightList []int
}

func main() {
	inputValues, err := GetInputValuesFromFile("day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	deviation := CalculateSimilarityScore(inputValues)
	fmt.Println("deviation: ", deviation)
}

func GetInputValuesFromFile(filePath string) (InputValues, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return InputValues{}, err
	}
	defer f.Close()

	result, err := GetInputValuesFromReader(f)
	if err != nil {
		return result, err
	}

	return result, nil
}

func GetInputValuesFromReader(reader io.Reader) (InputValues, error) {
	result := InputValues{LeftList: []int{}, RightList: []int{}}

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()

		var leftValue, rightValue int

		_, err := fmt.Sscanf(line, "%d %d", &leftValue, &rightValue)
		if err != nil {
			return result, err
		}

		result.LeftList = append(result.LeftList, leftValue)
		result.RightList = append(result.RightList, rightValue)
	}

	return result, nil
}

func CalculateSimilarityScore(inputValues InputValues) int {
	dict := make(map[int]int)

	for _, v := range inputValues.RightList {
		if _, ok := dict[v]; !ok {
			dict[v] = 1
		} else {
			dict[v] += 1
		}
	}

	sum := 0

	for _, v := range inputValues.LeftList {
		if _, ok := dict[v]; ok {
			sum += dict[v] * v
		}
	}

	return sum
}
