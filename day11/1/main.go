package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	stones, err := GetStonesFromFile("day11/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	newStones, err := PlayStones(stones, 25)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Len of stones: ", len(newStones))
}

func GetStonesFromFile(filename string) ([]int64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stones := make([]int64, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")

		for _, s := range split {
			n, err := strconv.ParseInt(s, 10, 64)
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

func PlayStones(stones []int64, n int) ([]int64, error) {
	currStones := make([]int64, len(stones))
	copy(currStones, stones)

	//fmt.Printf("Step 0, stones: %v\n", currStones)

	for i := 0; i < n; i++ {
		newStones := make([]int64, 0)

		for _, stone := range currStones {
			// make 1
			if stone == 0 {
				newStones = append(newStones, 1)
				continue
			}

			// split stones
			power := int64(1)
			length := 0
			for power <= stone {
				power *= 10
				length++
			}

			if length%2 == 0 {
				for range length / 2 {
					power /= 10
				}

				lStone := stone / power
				rStone := stone % power

				newStones = append(newStones, lStone)
				newStones = append(newStones, rStone)

				continue
			}

			// multiply the rest to 2024
			newStones = append(newStones, stone*2024)
		}

		currStones = newStones
		//fmt.Printf("Step %v, stones: %v\n", i+1, currStones)
	}

	return currStones, nil
}
