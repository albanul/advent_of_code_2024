package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/albanul/advent_of_code_2024/day10"
)

type Puzzle struct {
	Map           [][]int
	Width, Height int
}

type Point struct {
	i, j int
}

func main() {
	puzzle, err := GetPuzzleFromFile("day10/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum, err := GetSumOfAllTrailheadsScores(puzzle)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum: ", sum)
}

func GetPuzzleFromFile(filename string) (Puzzle, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Puzzle{}, err
	}
	defer file.Close()

	result := make([][]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := make([]int, 0)

		line := scanner.Text()
		for _, r := range line {
			n := int(r - '0')
			row = append(row, n)
		}

		result = append(result, row)
	}

	if err := scanner.Err(); err != nil {
		return Puzzle{}, err
	}

	puzzle := Puzzle{Map: result, Width: len(result[0]), Height: len(result)}
	return puzzle, nil
}

func GetSumOfAllTrailheadsScores(puzzle Puzzle) (int, error) {
	startingPoints := FindAllStartingPoints(puzzle)

	sum := 0

	for _, startingPoint := range startingPoints {
		score, err := GetTrailheadScore(puzzle, startingPoint)
		if err != nil {
			return sum, err
		}
		sum += score
	}

	return sum, nil
}

func FindAllStartingPoints(puzzle Puzzle) []Point {
	startingPoints := make([]Point, 0)

	for i := 0; i < puzzle.Height; i++ {
		for j := 0; j < puzzle.Width; j++ {
			if puzzle.Map[i][j] == 0 {
				startingPoints = append(startingPoints, Point{i, j})
			}
		}
	}

	return startingPoints
}

func GetTrailheadScore(puzzle Puzzle, startingPoint Point) (int, error) {
	historyMap := NewHistoryMap(puzzle)

	q := queue.NewQueue[Point]()
	q.Enqueue(startingPoint)

	for !q.IsEmpty() {
		p, err := q.Dequeue()
		if err != nil {
			return 0, err
		}

		historyMap[p.i][p.j] = 1

		pUp := Point{p.i - 1, p.j}
		pDown := Point{p.i + 1, p.j}
		pLeft := Point{p.i, p.j - 1}
		pRight := Point{p.i, p.j + 1}

		for _, pp := range []Point{pUp, pDown, pLeft, pRight} {
			if CanGoThere(puzzle, p, pp, historyMap) {
				q.Enqueue(pp)
			}
		}
	}

	score := 0
	for i := 0; i < puzzle.Height; i++ {
		for j := 0; j < puzzle.Width; j++ {
			if historyMap[i][j] == 1 && puzzle.Map[i][j] == 9 {
				score++
			}
		}
	}

	return score, nil
}

func NewHistoryMap(puzzle Puzzle) [][]int {
	historyMap := make([][]int, 0)
	for range puzzle.Height {
		line := make([]int, puzzle.Width)
		historyMap = append(historyMap, line)
	}
	return historyMap
}

func CanGoThere(puzzle Puzzle, p Point, pp Point, historyMap [][]int) bool {
	if pp.i < 0 || pp.j < 0 {
		return false
	}

	if pp.i >= puzzle.Height || pp.j >= puzzle.Width {
		return false
	}

	currValue := puzzle.Map[p.i][p.j]
	nextValue := puzzle.Map[pp.i][pp.j]

	if nextValue-currValue != 1 {
		return false
	}

	if historyMap[pp.i][pp.j] == 1 {
		return false
	}

	return true
}
