package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/albanul/advent_of_code_2024/day10"
)

type Point struct {
	i, j int
}

type Puzzle struct {
	Width, Height int
	Garden        [][]rune
}

func (p *Puzzle) GetValue(point Point) (rune, error) {
	if !WithinBorders(point, *p) {
		return rune(0), fmt.Errorf("point %v is outside borders", point)
	}

	return p.Garden[point.i][point.j], nil
}

func main() {
	puzzle, err := GetPuzzleFromFile("day12/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	price, err := CalculatePrice(puzzle)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Price: ", price)
}

func GetPuzzleFromFile(filepath string) (Puzzle, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return Puzzle{}, err
	}
	defer file.Close()

	garden := make([][]rune, 0)

	height := 0
	width := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		width = len(line)
		height++

		arr := make([]rune, 0)

		for _, r := range line {
			arr = append(arr, r)
		}

		garden = append(garden, arr)
	}

	if err := scanner.Err(); err != nil {
		return Puzzle{}, err
	}

	puzzle := Puzzle{width, height, garden}
	return puzzle, nil
}

func CalculatePrice(puzzle Puzzle) (int, error) {
	historyMap := make(map[Point]bool)

	regionQueue := queue.NewQueue[Point]()
	areaQueue := queue.NewQueue[Point]()

	areaQueue.Enqueue(Point{0, 0})

	sum := 0

	for !areaQueue.IsEmpty() {
		firstRegionPoint, err := areaQueue.Dequeue()
		if err != nil {
			return 0, err
		}

		area, perimeter := 0, 0

		regionQueue.Enqueue(firstRegionPoint)

		for !regionQueue.IsEmpty() {
			currPoint, err := regionQueue.Dequeue()
			if err != nil {
				return 0, err
			}

			currentValue, err := puzzle.GetValue(currPoint)
			if err != nil {
				return 0, err
			}

			if _, wasThere := historyMap[currPoint]; wasThere {
				continue
			}
			historyMap[currPoint] = true
			area++

			up := Point{currPoint.i - 1, currPoint.j}
			down := Point{currPoint.i + 1, currPoint.j}
			left := Point{currPoint.i, currPoint.j - 1}
			right := Point{currPoint.i, currPoint.j + 1}

			pointPerimeter := 0

			for _, p := range []Point{up, down, left, right} {
				if !WithinBorders(p, puzzle) {
					pointPerimeter++
					continue
				}

				pv, err := puzzle.GetValue(p)
				if err != nil {
					return 0, err
				}

				if pv != currentValue {
					pointPerimeter++
					areaQueue.Enqueue(p)
				} else {
					regionQueue.Enqueue(p)
				}
			}

			perimeter += pointPerimeter
		}

		sum += area * perimeter
	}

	return sum, nil
}

func WithinBorders(point Point, puzzle Puzzle) bool {
	if point.i < 0 || point.j < 0 || point.i >= puzzle.Width || point.j >= puzzle.Height {
		return false
	}

	return true
}
