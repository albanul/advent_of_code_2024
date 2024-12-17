package day15

import (
	"errors"
	"fmt"
)

type Point struct {
	I, J int
}

func (p *Point) GetGPS() int {
	return p.I*100 + p.J
}

type MoveDirection int

const (
	MoveDirectionUp MoveDirection = iota
	MoveDirectionRight
	MoveDirectionDown
	MoveDirectionLeft
)

type Game struct {
	walls         map[Point]bool
	boxes         map[Point]bool
	robotPosition Point
	robotMoves    []MoveDirection
	width, height int
}

func (g *Game) PlayGame(drawSteps bool) error {
	if drawSteps {
		g.Draw()
	}

	for _, move := range g.robotMoves {
		err := g.moveRobot(move)
		if err != nil {
			return err
		}

		if drawSteps {
			g.Draw()
		}
	}

	return nil
}

func (g *Game) CalculateGPSSum() int {
	sum := 0
	for p := range g.boxes {
		sum += p.GetGPS()
	}

	return sum
}

func (g *Game) moveRobot(move MoveDirection) error {
	di, dj, err := toDeltas(move)
	if err != nil {
		return err
	}

	nP := Point{g.robotPosition.I + di, g.robotPosition.J + dj}

	// if it's a wall then don't move
	if _, isWall := g.walls[nP]; isWall {
		return nil
	}

	tP := nP

	boxesToMove := make([]Point, 0)

	_, isBox := g.boxes[tP]

	// if it's not a wall and not a box, then just move there
	if !isBox {
		g.robotPosition = nP
		return nil
	}

	// move boxes
	for isBox {
		boxesToMove = append(boxesToMove, tP)
		tP = Point{tP.I + di, tP.J + dj}

		_, isBox = g.boxes[tP]
		_, isWall := g.walls[tP]

		if isWall {
			boxesToMove = make([]Point, 0)
			break
		}
	}

	// if no boxes to move, then don't move
	if len(boxesToMove) == 0 {
		return nil
	}

	delete(g.boxes, boxesToMove[0])

	for _, p := range boxesToMove {
		nP := Point{p.I + di, p.J + dj}
		g.boxes[nP] = true
	}

	g.robotPosition = nP

	return nil
}

func (g *Game) Draw() {
	fmt.Println()
	for i := 0; i < g.height; i++ {
		for j := 0; j < g.width; j++ {
			if g.walls[Point{i, j}] {
				fmt.Print("#")
				continue
			}

			if g.boxes[Point{i, j}] {
				fmt.Print("O")
				continue
			}

			p := Point{i, j}
			if p == g.robotPosition {
				fmt.Print("@")
				continue
			}

			fmt.Print(".")
		}
		fmt.Println()
	}
}

func NewGame(walls, boxes map[Point]bool, robotPosition Point, robotMoves []MoveDirection, width, height int) *Game {
	game := &Game{
		walls:         walls,
		boxes:         boxes,
		robotPosition: robotPosition,
		robotMoves:    robotMoves,
		width:         width,
		height:        height,
	}

	return game
}

func toDeltas(direction MoveDirection) (di, dj int, err error) {
	switch direction {
	case MoveDirectionUp:
		return -1, 0, nil
	case MoveDirectionRight:
		return 0, 1, nil
	case MoveDirectionDown:
		return 1, 0, nil
	case MoveDirectionLeft:
		return 0, -1, nil
	default:
		return 0, 0, errors.New("invalid direction")
	}
}
