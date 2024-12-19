package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/albanul/advent_of_code_2024/day18/2/game"
)

func main() {
	walls, err := getWallsFromFile("day18/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var p day18.Point

	for time := 1024; time < len(walls); time++ {
		game := day18.NewGame(walls, time, 71, 71)
		result, err := game.Play()
		if err != nil {
			log.Fatal(err)
		}

		if result == -1 {
			fmt.Println("The game is over. Time:", time)
			p = walls[time-1]
			break
		}
	}

	fmt.Printf("Point: %d,%d", p.I, p.J)
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
