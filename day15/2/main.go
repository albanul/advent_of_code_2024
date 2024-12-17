package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/albanul/advent_of_code_2024/day15/game_v2"
)

// 1548028 - incorrect
func main() {
	game, err := getGameFromFile("day15/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	err = game.PlayGame(false)
	if err != nil {
		log.Fatal(err)
	}

	//game.Draw()

	sum := game.CalculateGPSSum()
	fmt.Println("sum: ", sum)
}

func getGameFromFile(filename string) (*day15.Game, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	walls := make(map[day15.Point]bool)
	boxes := make(map[day15.Point]*day15.Box)
	robotPosition := day15.Point{I: 0, J: 0}
	robotMoves := make([]day15.MoveDirection, 0)

	i, width, height := 0, 0, 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		for j, v := range line {
			point := day15.Point{I: i, J: 2 * j}
			point2 := point
			point2.J += 1

			if v == '#' {
				walls[point] = true
				walls[point2] = true
			}

			if v == 'O' {
				box := day15.NewBox(point)

				boxes[point] = box
				boxes[point2] = box
			}

			if v == '@' {
				robotPosition = point
			}
		}

		width = 2 * len(line)
		i++
	}
	height = i

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	for scanner.Scan() {
		line := scanner.Text()

		for _, v := range line {
			if v == '^' {
				robotMoves = append(robotMoves, day15.MoveDirectionUp)
			}
			if v == '>' {
				robotMoves = append(robotMoves, day15.MoveDirectionRight)
			}
			if v == 'v' {
				robotMoves = append(robotMoves, day15.MoveDirectionDown)
			}
			if v == '<' {
				robotMoves = append(robotMoves, day15.MoveDirectionLeft)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	game := day15.NewGame(walls, boxes, robotPosition, robotMoves, width, height)
	return game, nil
}
