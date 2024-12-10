package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

type Direction int

const (
	Top Direction = iota
	Right
	Bottom
	Left
)

type Position struct {
	X, Y int
}

type Guard struct {
	Position  Position
	Direction Direction
}

type Game struct {
	width, height int
	counter       int
	patrolPath    map[Position]bool
	walls         map[Position]bool
	guard         Guard
}

func (g *Game) SetWall(x, y int) {
	position := Position{x, y}
	g.walls[position] = true
}

func (g *Game) SetGuard(guard Guard) {
	g.counter = 1
	g.guard = guard
	g.patrolPath[g.guard.Position] = true
}

func (g *Game) MakeMove() {
	deltaX, deltaY := 0, 0
	if g.guard.Direction == Top {
		deltaX, deltaY = 0, -1
	}

	if g.guard.Direction == Bottom {
		deltaX, deltaY = 0, 1
	}

	if g.guard.Direction == Left {
		deltaX, deltaY = -1, 0
	}

	if g.guard.Direction == Right {
		deltaX, deltaY = 1, 0
	}

	newPosition := Position{g.guard.Position.X + deltaX, g.guard.Position.Y + deltaY}
	isWall := g.walls[newPosition]

	if isWall {
		newDirection := (g.guard.Direction + 1) % 4
		g.guard.Direction = newDirection
	} else {
		wasVisited := g.patrolPath[newPosition]
		if !wasVisited {
			g.patrolPath[newPosition] = true
			g.counter++
		}
		g.guard.Position = newPosition
	}
}

func (g *Game) GetCounter() int {
	return g.counter - 1
}

func (g *Game) IsOver() bool {
	return g.guard.Position.X >= g.width || g.guard.Position.X < 0 ||
		g.guard.Position.Y >= g.height || g.guard.Position.Y < 0
}

func main() {
	game, err := GetGameFromFile("day06/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	PlayGame(game)
	counter := game.GetCounter()
	log.Println("Counter:", counter)
}

func GetGameFromFile(filepath string) (*Game, error) {
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

func GetGame(file io.Reader) (*Game, error) {
	scanner := bufio.NewScanner(file)

	game := &Game{walls: make(map[Position]bool), counter: 0, patrolPath: make(map[Position]bool)}

	y := 0

	for scanner.Scan() {
		line := scanner.Text()
		game.width = len(line)

		for x, ch := range line {
			if ch == '#' {
				game.SetWall(x, y)
			} else if ch == '^' {
				guard := Guard{Position: Position{x, y}, Direction: Top}
				game.SetGuard(guard)
			} else if ch == '<' {
				guard := Guard{Position: Position{x, y}, Direction: Left}
				game.SetGuard(guard)
			} else if ch == '>' {
				guard := Guard{Position: Position{x, y}, Direction: Right}
				game.SetGuard(guard)
			} else if ch == 'v' {
				guard := Guard{Position: Position{x, y}, Direction: Bottom}
				game.SetGuard(guard)
			}
		}
		y++
	}

	game.height = y

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return game, nil
}

func PlayGame(game *Game) {
	for !game.IsOver() {
		game.MakeMove()
	}
}
