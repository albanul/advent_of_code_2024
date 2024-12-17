package day15

import (
	"errors"
	"fmt"

	queue "github.com/albanul/advent_of_code_2024/day10"
)

type Point struct {
	I, J int
}

func (p *Point) GetGPS() int {
	return p.I*100 + p.J
}

type Box struct {
	leftPosition Point
}

func NewBox(point Point) *Box {
	return &Box{leftPosition: point}
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
	boxes         map[Point]*Box
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
			drawMove(move)
			g.Draw()
		}
	}

	return nil
}

func (g *Game) CalculateGPSSum() int {
	sum := 0

	for i := 0; i < g.height; i++ {
		for j := 0; j < g.width; j++ {
			p := Point{i, j}

			if _, ok := g.boxes[p]; ok {
				sum += p.GetGPS()
				j++
			}
		}
	}
	return sum
}

func (g *Game) Draw() {
	for i := 0; i < g.height; i++ {
		for j := 0; j < g.width; j++ {
			if g.walls[Point{i, j}] {
				fmt.Print("#")
				continue
			}

			if g.boxes[Point{i, j}] != nil {
				fmt.Print("[]")
				j++
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
	fmt.Println()
}

func (g *Game) moveRobot(move MoveDirection) error {
	di, dj, err := toDeltas(move)
	if err != nil {
		return err
	}

	nextPoint := Point{g.robotPosition.I + di, g.robotPosition.J + dj}

	// if it's a wall then don't move
	if _, isWall := g.walls[nextPoint]; isWall {
		return nil
	}

	box, isBox := g.boxes[nextPoint]

	// if it's not a wall and not a box, then just move there
	if !isBox {
		g.robotPosition = nextPoint
		return nil
	}

	// move boxes
	boxesToMove := make([]*Box, 0)

	q := queue.NewQueue[*Box]()
	q.Enqueue(box)

	for !q.IsEmpty() {
		box, err = q.Dequeue()
		if err != nil {
			return err
		}

		boxesToMove = append(boxesToMove, box)

		neighbours, nextToWall, err := g.findNeighbours(box, move)
		if err != nil {
			return err
		}

		for _, n := range neighbours {
			q.Enqueue(n)
		}

		// if something is adjusted to the wall then don't move any boxes
		if nextToWall {
			boxesToMove = boxesToMove[:0]
			break
		}
	}

	// if no boxes to move, then don't move
	if len(boxesToMove) == 0 {
		return nil
	}

	for i := len(boxesToMove) - 1; i >= 0; i-- {
		box := boxesToMove[i]

		g.removeBox(box)

		box.leftPosition.I += di
		box.leftPosition.J += dj

		g.addBox(box)
	}

	g.robotPosition = nextPoint

	return nil
}

func (g *Game) removeBox(box *Box) {
	p1 := box.leftPosition
	p2 := p1
	p2.J += 1

	delete(g.boxes, p1)
	delete(g.boxes, p2)
}

func (g *Game) addBox(box *Box) {
	p1 := box.leftPosition
	p2 := p1
	p2.J += 1

	g.boxes[p1] = box
	g.boxes[p2] = box
}

func (g *Game) findNeighbours(box *Box, direction MoveDirection) ([]*Box, bool, error) {
	nextToWall := false
	di, _, err := toDeltas(direction)
	if err != nil {
		return nil, false, err
	}

	neighbours := make([]*Box, 0)

	switch direction {
	case MoveDirectionUp:
		fallthrough
	case MoveDirectionDown:
		p1 := box.leftPosition
		p1.I += di
		p2 := p1
		p2.J += 1

		b1, ok := g.boxes[p1]
		if ok {
			neighbours = append(neighbours, b1)
		}

		if _, ok := g.walls[p1]; ok {
			nextToWall = true
		}

		b2, ok := g.boxes[p2]
		if ok && b1 != b2 {
			neighbours = append(neighbours, b2)
		}

		if _, ok := g.walls[p2]; ok {
			nextToWall = true
		}
	case MoveDirectionLeft:
		p := box.leftPosition
		p.J -= 1

		if b, ok := g.boxes[p]; ok {
			neighbours = append(neighbours, b)
		}

		if _, ok := g.walls[p]; ok {
			nextToWall = true
		}
	case MoveDirectionRight:
		p := box.leftPosition
		p.J += 2

		if b, ok := g.boxes[p]; ok {
			neighbours = append(neighbours, b)
		}

		if _, ok := g.walls[p]; ok {
			nextToWall = true
		}
	default:
		return nil, false, errors.New("invalid direction")
	}

	return neighbours, nextToWall, nil
}

func NewGame(walls map[Point]bool, boxes map[Point]*Box, robotPosition Point, robotMoves []MoveDirection, width, height int) *Game {
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

func drawMove(move MoveDirection) {
	var r string
	if move == MoveDirectionUp {
		r = "^"
	}
	if move == MoveDirectionRight {
		r = ">"
	}
	if move == MoveDirectionDown {
		r = "v"
	}
	if move == MoveDirectionLeft {
		r = "<"
	}

	fmt.Println(r)
}
