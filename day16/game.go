package day16

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/albanul/advent_of_code_2024/day10"
)

type MoveDirection int

const (
	MoveDirectionUp MoveDirection = iota
	MoveDirectionRight
	MoveDirectionDown
	MoveDirectionLeft
)

type Point struct {
	I, J int
}

func (p Point) NextPoint(direction MoveDirection) (Point, error) {
	di, dj, err := toDeltas(direction)
	if err != nil {
		return Point{}, err
	}

	newPoint := p
	newPoint.I += di
	newPoint.J += dj

	return newPoint, nil
}

type turn struct {
	score     int
	path      []Point
	point     Point
	direction MoveDirection
}

type Game struct {
	start, end    Point
	weightMap     [][]int
	paths         map[Point][][]Point
	walls         map[Point]bool
	width, height int
}

func (g *Game) Play() error {
	q := queue.NewQueue[turn]()

	pd := turn{point: g.start, score: 0, path: []Point{g.start}, direction: MoveDirectionRight}
	q.Enqueue(pd)

	weightIncreases := [4]int{1, 1001, 2001, 1001}

	for !q.IsEmpty() {
		currTurn, err := q.Dequeue()
		if err != nil {
			return err
		}
		p, dir, s, path := currTurn.point, currTurn.direction, currTurn.score, currTurn.path

		score, err := g.getScore(p)
		if err != nil {
			return err
		}

		if score == -1 || s < score {
			g.weightMap[p.I][p.J] = s
			g.paths[p] = [][]Point{path}
		} else if s == score {
			g.paths[p] = append(g.paths[p], path)
		}

		for i, wInc := range weightIncreases {
			newDirection := MoveDirection((int(dir) + i) % 4)

			newPoint, err := p.NextPoint(newDirection)
			if err != nil {
				return err
			}

			if g.canGoThere(newPoint) {
				score, err := g.getScore(newPoint)
				if err != nil {
					return err
				}

				newScore := s + wInc

				newPath := make([]Point, len(path))
				copy(newPath, path)
				newPath = append(newPath, newPoint)

				if score == -1 || newScore < score {
					nPD := turn{score: newScore, path: newPath, point: newPoint, direction: newDirection}
					q.Enqueue(nPD)
				}
			}
		}
	}

	return nil
}

func (g *Game) Draw() {
	for i := 0; i < g.height; i++ {
		for j := 0; j < g.width; j++ {
			p := Point{i, j}
			if _, ok := g.walls[p]; ok {
				fmt.Print("#")
				continue
			}

			if p == g.start {
				fmt.Print("S")
				continue
			}

			if p == g.end {
				fmt.Print("E")
				continue
			}

			fmt.Print(strconv.Itoa(g.weightMap[p.I][p.J]))
		}
		fmt.Println()
	}
}

func (g *Game) GetFinalScore() (int, error) {
	score := g.weightMap[g.end.I][g.end.J]
	if score == -1 {
		return -1, errors.New("there is no path from start point to end point")
	}

	return score, nil
}

func (g *Game) GetTilesCount(draw bool) (int, error) {
	paths := g.paths[g.end]

	if len(paths) == 0 {
		return 0, errors.New("there is no path from start point to end point")
	}

	visited := make(map[Point]bool)
	for _, path := range paths {
		for _, point := range path {
			visited[point] = true
		}
	}

	if draw {
		fmt.Println("All visited points:")
		for i := 0; i < g.height; i++ {
			for j := 0; j < g.width; j++ {
				p := Point{I: i, J: j}
				if _, ok := visited[p]; ok {
					fmt.Print("O")
					continue
				}
				if _, ok := g.walls[p]; ok {
					fmt.Print("#")
					continue
				}
				fmt.Print(".")
			}
			fmt.Println()
		}
	}

	return len(visited), nil
}

func (g *Game) canGoThere(p Point) bool {
	if p.I < 0 || p.J < 0 || p.I >= g.height || p.J >= g.width {
		return false
	}

	if _, ok := g.walls[p]; ok {
		return false
	}

	return true
}

func (g *Game) getScore(p Point) (int, error) {
	if !g.canGoThere(p) {
		return 0, errors.New(fmt.Sprintf("invalid point (%d,%d)", p.I, p.J))
	}
	return g.weightMap[p.I][p.J], nil
}

func NewGame(width, height int, start, end Point, walls map[Point]bool) *Game {
	weightMap := make([][]int, height)
	for i := range weightMap {
		weightMap[i] = make([]int, 0)
		for range width {
			weightMap[i] = append(weightMap[i], -1)
		}
	}

	paths := make(map[Point][][]Point)

	weightMap[start.I][start.J] = 0

	g := &Game{
		start:     start,
		end:       end,
		walls:     walls,
		width:     width,
		height:    height,
		weightMap: weightMap,
		paths:     paths,
	}

	return g
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
