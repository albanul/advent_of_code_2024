package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/albanul/advent_of_code_2024/day14"
)

func main() {
	robots, err := GetRobotsFromFile("day14/input.txt")
	if err != nil {
		panic(err)
	}

	game := day14.NewGame(101, 103, robots)
	game.PlayFor(100)
	sf := day14.CalculateSafetyFactor(game)

	fmt.Println("Safety factor: ", sf)
}

func GetRobotsFromFile(file string) ([]*day14.Robot, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	robots := make([]*day14.Robot, 0)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		var i, j, di, dj int

		_, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &j, &i, &dj, &di)
		if err != nil {
			return nil, err
		}

		point := day14.Point{I: i, J: j}
		velocity := day14.Velocity{Di: di, Dj: dj}
		robot := &day14.Robot{Position: point, Velocity: velocity}

		robots = append(robots, robot)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return robots, nil
}
