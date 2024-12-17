package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/albanul/advent_of_code_2024/day16"
)

func main() {
	game, err := GetGameFromFile("day16/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	//game.Draw()

	err = game.Play()
	if err != nil {
		log.Fatal(err)
	}

	//game.Draw()

	score, err := game.GetFinalScore()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("score: ", score)

}

func GetGameFromFile(filepath string) (*day16.Game, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var start, end day16.Point
	walls := make(map[day16.Point]bool)
	width, height := 0, 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		for j, v := range line {
			i := height
			p := day16.Point{I: i, J: j}
			if v == '#' {
				walls[p] = true
				continue
			}

			if v == 'S' {
				start = p
				continue
			}

			if v == 'E' {
				end = p
				continue
			}
		}

		width = len(line)
		height++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	game := day16.NewGame(width, height, start, end, walls)
	return game, nil
}
