package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/albanul/advent_of_code_2024/day17"
)

func main() {
	game, err := getGameFromFile("day17/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	err = game.Play()
	if err != nil {
		log.Fatal(err)
	}

	output := game.GetOutput()
	fmt.Println("output:", output)
}

func getGameFromFile(filename string) (*day17.Game, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	regArray := [3]int{}
	input := make([]int, 0)

	scanner := bufio.NewScanner(f)

	for i := range 3 {
		scanner.Scan()
		line := scanner.Text()

		var regName string
		_, err := fmt.Sscanf(line, "Register %s %d", &regName, &regArray[i])
		if err != nil {
			return nil, err
		}
	}

	scanner.Scan()
	scanner.Scan()
	line := scanner.Text()

	split := strings.Split(line, ":")
	inputString := split[1]
	inputString = strings.TrimLeft(inputString, " ")
	split = strings.Split(inputString, ",")

	for s := range split {
		n, err := strconv.Atoi(split[s])
		if err != nil {
			return nil, err
		}

		input = append(input, int(n))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	game := day17.NewGame(regArray[0], regArray[1], regArray[2], input)

	return game, nil
}
