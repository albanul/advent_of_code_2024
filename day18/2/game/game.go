package day18

import (
	"errors"
	"fmt"

	queue "github.com/albanul/advent_of_code_2024/day10"
)

type Point struct {
	I, J int
}

type Game struct {
	weightMap     [][]int
	walls         map[Point]bool
	width, height int
}

func (g *Game) Play() (int, error) {
	q := queue.NewQueue[Point]()

	start := Point{0, 0}
	g.weightMap[0][0] = 1
	q.Enqueue(start)

	for !q.IsEmpty() {
		p, err := q.Dequeue()
		if err != nil {
			return 0, err
		}

		score, err := g.getWeight(p)
		if err != nil {
			return 0, err
		}

		up := Point{p.I - 1, p.J}
		down := Point{p.I + 1, p.J}
		left := Point{p.I, p.J - 1}
		right := Point{p.I, p.J + 1}

		directions := []Point{up, down, left, right}

		for _, newPoint := range directions {
			if g.canGoThere(newPoint) {
				newScore := score + 1

				score, err := g.getWeight(newPoint)
				if err != nil {
					return 0, err
				}

				if score == 0 || newScore < score {
					g.weightMap[newPoint.I][newPoint.J] = newScore
					q.Enqueue(newPoint)
				}
			}
		}
	}

	end := Point{g.height - 1, g.width - 1}
	result, err := g.getWeight(end)
	if err != nil {
		return 0, err
	}

	return result - 1, nil
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

func (g *Game) getWeight(p Point) (int, error) {
	if !g.canGoThere(p) {
		return 0, errors.New(fmt.Sprintf("invalid point (%d,%d)", p.I, p.J))
	}
	return g.weightMap[p.I][p.J], nil
}

func NewGame(allWalls []Point, time, width, height int) *Game {
	walls := make(map[Point]bool)
	for i := range time {
		p := allWalls[i]
		walls[p] = true
	}

	weightMap := make([][]int, height)
	for i := range height {
		weightMap[i] = make([]int, width)
	}

	return &Game{weightMap, walls, width, height}
}
