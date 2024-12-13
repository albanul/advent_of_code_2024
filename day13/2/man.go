package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

type Button struct {
	Cost   int
	Dx, Dy int
}

type Point struct {
	X, Y int64
}

type Game struct {
	Buttons []Button
	Prize   Point
}

func main() {
	games, err := GetGamesFromFile("day13/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum := CalculateSumOfTokens(games)
	fmt.Println("sum: ", sum)
}

func GetGamesFromFile(file string) ([]Game, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var games []Game

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		buttonA := Button{Cost: 3}
		buttonB := Button{Cost: 1}

		lineA := scanner.Text()
		_, err := fmt.Sscanf(lineA, "Button A: X+%d, Y+%d", &buttonA.Dx, &buttonA.Dy)
		if err != nil {
			return nil, err
		}

		scanner.Scan()
		lineB := scanner.Text()
		_, err = fmt.Sscanf(lineB, "Button B: X+%d, Y+%d", &buttonB.Dx, &buttonB.Dy)
		if err != nil {
			return nil, err
		}

		point := Point{X: 0, Y: 0}

		scanner.Scan()
		priceLine := scanner.Text()
		_, err = fmt.Sscanf(priceLine, "Prize: X=%d, Y=%d", &point.X, &point.Y)
		if err != nil {
			return nil, err
		}

		game := Game{[]Button{buttonA, buttonB}, point}
		games = append(games, game)

		if !scanner.Scan() {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return games, nil
}

func CalculateSumOfTokens(games []Game) int64 {
	sum := int64(0)

	for _, game := range games {
		a, b, err := TryToWin(game)

		if err != nil {
			continue
		}

		sum += int64(3*a) + int64(b)
	}

	return sum
}

func TryToWin(game Game) (a, b int, err error) {
	b1 := game.Buttons[0]
	b2 := game.Buttons[1]

	add := int64(10_000_000_000_000)

	fb := float64((game.Prize.Y+add)*int64(b1.Dx)-(game.Prize.X+add)*int64(b1.Dy)) / float64(b1.Dx*b2.Dy-b2.Dx*b1.Dy)
	b = int(fb)

	if fb != float64(b) {
		return a, b, errors.New("can't win")
	}

	fa := float64((game.Prize.X+add)-int64(b*b2.Dx)) / float64(b1.Dx)
	a = int(fa)
	if fa != float64(a) {
		return a, b, errors.New("can't win")
	}

	return a, b, nil
}
