package main

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/albanul/advent_of_code_2024/day06"
)

func main() {
	game, err := GetGameFromFile("day06/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	counter := GetObstructionsCounter(game)
	log.Println("Counter:", counter)
}

func GetGameFromFile(filepath string) (*day06.Game, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	game, err := GetGame(file)
	if err != nil {
		return nil, err
	}

	return game, nil
}

func GetGame(file io.Reader) (*day06.Game, error) {
	scanner := bufio.NewScanner(file)

	x, y := 0, 0
	walls := make(map[day06.Position]bool)
	var guard day06.Guard

	for scanner.Scan() {
		line := scanner.Text()
		x = len(line)

		for x, ch := range line {
			if ch == '#' {
				walls[day06.Position{X: x, Y: y}] = true
			} else if ch == '^' {
				guard = day06.Guard{Position: day06.Position{X: x, Y: y}, Direction: day06.Top}
			} else if ch == '<' {
				guard = day06.Guard{Position: day06.Position{X: x, Y: y}, Direction: day06.Left}
			} else if ch == '>' {
				guard = day06.Guard{Position: day06.Position{X: x, Y: y}, Direction: day06.Right}
			} else if ch == 'v' {
				guard = day06.Guard{Position: day06.Position{X: x, Y: y}, Direction: day06.Bottom}
			}
		}
		y++
	}

	game := day06.NewGame(x, y)
	game.SetWalls(walls)
	game.SetGuard(guard)

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return game, nil
}

func GetObstructionsCounter(game *day06.Game) int {
	counter := 0

	x, y := game.GetSizes()

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if !game.CanSetWall(j, i) {
				continue
			}

			game.SetWall(j, i)
			PlayGame(game)

			if game.IsInLoop() {
				counter++
			}

			wallPosition := day06.Position{X: j, Y: i}
			game.Reset(wallPosition)
		}
	}

	return counter
}

func PlayGame(game *day06.Game) {
	for !game.IsOver() && !game.IsInLoop() {
		game.MakeMove()
	}
}
