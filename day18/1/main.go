package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/albanul/advent_of_code_2024/day18/1/game"
)

func main() {
	walls, err := getWallsFromFile("day18/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	game := day18.NewGame(walls, 1024, 71, 71)
	result, err := game.Play()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result:", result)
}

func getWallsFromFile(filename string) ([]day18.Point, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	allWalls := make([]day18.Point, 0)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, ",")
		j, err := strconv.Atoi(split[0])
		if err != nil {
			return nil, err
		}
		i, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, err
		}

		wall := day18.Point{I: j, J: i}
		allWalls = append(allWalls, wall)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return allWalls, nil
}
