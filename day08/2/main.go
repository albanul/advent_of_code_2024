package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

type Point struct {
	i, j int
}

type Puzzle struct {
	Width, Height int
	Antennas      map[rune][]Point
}

func main() {
	puzzle, err := GetPuzzleFromFile("day08/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	count := CountUniqueAntinodes(puzzle)
	fmt.Println("Count: ", count)
}

func GetPuzzleFromFile(filepath string) (*Puzzle, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, errors.Join(errors.New("can't open file "+filepath), err)
	}
	defer file.Close()

	puzzle, err := ParsePuzzle(file)
	if err != nil {
		return nil, errors.Join(errors.New("can't parse file "+filepath), err)
	}

	return puzzle, nil
}

func ParsePuzzle(reader io.Reader) (*Puzzle, error) {
	puzzle := new(Puzzle)
	scanner := bufio.NewScanner(reader)

	width, height := 0, 0

	antennas := make(map[rune][]Point)

	for scanner.Scan() {
		line := scanner.Text()

		for j, ch := range line {
			if unicode.IsLetter(ch) || unicode.IsNumber(ch) {
				_, ok := antennas[ch]
				if !ok {
					antennas[ch] = make([]Point, 0)
				}

				coordinate := Point{height, j}
				antennas[ch] = append(antennas[ch], coordinate)
			}
		}

		width = len(line)
		height++
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.Join(errors.New("something went wrong during file parsing"), err)
	}

	puzzle.Width = width
	puzzle.Height = height
	puzzle.Antennas = antennas

	return puzzle, nil
}

func CountUniqueAntinodes(puzzle *Puzzle) int {
	m := make(map[Point]bool)

	for _, coordinates := range puzzle.Antennas {
		for i := 0; i < len(coordinates)-1; i++ {
			x := coordinates[i]
			for j := i + 1; j < len(coordinates); j++ {
				y := coordinates[j]

				di := x.i - y.i
				dj := x.j - y.j

				antinode1 := Point{x.i, x.j}
				for IsPointInsideArea(antinode1, puzzle.Width, puzzle.Height) {
					m[antinode1] = true
					antinode1.i = antinode1.i + di
					antinode1.j = antinode1.j + dj
				}

				antinode2 := Point{y.i, y.j}
				for IsPointInsideArea(antinode2, puzzle.Width, puzzle.Height) {
					m[antinode2] = true
					antinode2.i = antinode2.i - di
					antinode2.j = antinode2.j - dj
				}
			}
		}
	}

	return len(m)
}

func IsPointInsideArea(p Point, width, height int) bool {
	return p.i >= 0 && p.j >= 0 && p.i < width && p.j < height
}
