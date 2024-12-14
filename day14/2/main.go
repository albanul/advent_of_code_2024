package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/albanul/advent_of_code_2024/day14"
)

func main() {
	robots, err := GetRobotsFromFile("day14/input.txt")
	if err != nil {
		panic(err)
	}

	game := day14.NewGame(101, 103, robots)

	rr := game.GetRobots()
	m := GetRobotsMap(rr)
	DrawGame(game, 0, m)

	// I don't have any other solution how to find a Christmas tree rather than
	// have a look on all the generated pictures and find the Christmas tree yourself
	// ¯\_(ツ)_/¯
	for second := range 10000 {
		game.PlayFor1Second()

		if second < 5000 {
			continue
		}

		robots := game.GetRobots()

		m := GetRobotsMap(robots)
		DrawGame(game, second+1, m)
	}

	sf := day14.CalculateSafetyFactor(game)

	fmt.Println("Safety factor: ", sf)
}

func GetRobotsMap(robots []day14.Robot) map[day14.Point]bool {
	m := make(map[day14.Point]bool)
	for _, r := range robots {
		if _, ok := m[r.Position]; !ok {
			m[r.Position] = true
		}
	}
	return m
}

func DrawGame(game *day14.Game, second int, m map[day14.Point]bool) {
	pixels := make([][]int, 0)

	for i := 0; i < game.GetHeight(); i++ {
		pixels = append(pixels, make([]int, game.GetWidth()))

		for j := 0; j < game.GetWidth(); j++ {
			p := day14.Point{I: i, J: j}
			if _, ok := m[p]; ok {
				pixels[i][j] = 1
			}
		}
	}

	err := createPNG(game.GetWidth(), game.GetHeight(), pixels, fmt.Sprintf("day14/output/%d.png", second))
	if err != nil {
		log.Fatal(err)
	}
}

func GetRobotsFromFile(file string) ([]*day14.Robot, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	robots := make([]*day14.Robot, 0)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		var i, j, di, dj int

		_, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &j, &i, &dj, &di)
		if err != nil {
			return nil, err
		}

		point := day14.Point{I: i, J: j}
		velocity := day14.Velocity{Di: di, Dj: dj}
		robot := &day14.Robot{Position: point, Velocity: velocity}

		robots = append(robots, robot)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return robots, nil
}

func createPNG(width, height int, pixelData [][]int, outputFile string) error {
	// Create a new blank image
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Set pixels based on the input pixelData
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if pixelData[y][x] == 1 {
				img.Set(x, y, color.White) // Set pixel to white
			} else {
				img.Set(x, y, color.Black) // Set pixel to black
			}
		}
	}

	// Create and open the output file
	file, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode the image to PNG and write to the file
	if err := png.Encode(file, img); err != nil {
		return err
	}

	return nil
}
