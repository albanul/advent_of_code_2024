package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"slices"
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

	deviation := CalculateDeviation(inputValues)
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

func CalculateDeviation(inputValues InputValues) int {
	slices.Sort(inputValues.LeftList)
	slices.Sort(inputValues.RightList)

	sum := 0

	for i := 0; i < len(inputValues.LeftList); i++ {
		deviationWithSign := inputValues.LeftList[i] - inputValues.RightList[i]
		sum += int(math.Abs(float64(deviationWithSign)))
	}

	return sum
}
