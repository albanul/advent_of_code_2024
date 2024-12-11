package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Stone struct {
	value       int
	coefficient int
}

func main() {
	stones, err := GetStonesFromFile("day11/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	nStones, err := PlayStones(stones, 75)
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	for _, v := range nStones {
		sum += v.coefficient
	}

	fmt.Println("Count of stones: ", sum)
}

func GetStonesFromFile(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stones := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")

		for _, s := range split {
			n, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}

			stones = append(stones, n)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return stones, nil
}

func PlayStones(stones []int, n int) ([]Stone, error) {
	newStones := make([]Stone, 0)
	for _, stone := range stones {
		newStones = append(newStones, Stone{stone, 1})
	}

	for i := 0; i < n; i++ {
		newStones = PlayOneTurn(newStones)
	}

	return newStones, nil
}

func PlayOneTurn(stones []Stone) []Stone {
	m := make(map[int]int)

	for _, stone := range stones {
		// make 1 from 0
		if stone.value == 0 {
			newStone := Stone{1, stone.coefficient}
			AddToMap(m, newStone)
			continue
		}

		// split stones
		length := int(math.Log10(float64(stone.value))) + 1

		if length%2 == 0 {
			divisor := int(math.Pow10(length / 2))

			lStone := Stone{stone.value / divisor, stone.coefficient}
			rStone := Stone{stone.value % divisor, stone.coefficient}

			AddToMap(m, lStone)
			AddToMap(m, rStone)

			continue
		}

		// multiply the rest to 2024
		newStone := Stone{stone.value * 2024, stone.coefficient}
		AddToMap(m, newStone)
	}

	newStones := make([]Stone, 0)
	for k, v := range m {
		newStones = append(newStones, Stone{k, v})
	}
	return newStones
}

func AddToMap(m map[int]int, stone Stone) {
	c, ok := m[stone.value]
	if !ok {
		m[stone.value] = stone.coefficient
	}
	m[stone.value] = c + stone.coefficient
}
