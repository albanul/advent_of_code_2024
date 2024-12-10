package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Direction int

const (
	DirectionUnknown Direction = iota
	DirectionUp
	DirectionDown
)

type Report []int

func main() {
	reports, err := GetReportsFromFile("day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	count := CountSafeReports(reports)
	fmt.Println("Safe reports: ", count)
}

func GetReportsFromFile(filename string) ([]Report, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reports, err := GetReports(file)
	if err != nil {
		return nil, err
	}

	return reports, nil
}

func GetReports(file io.Reader) ([]Report, error) {
	reports := make([]Report, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		report, err := ScanReportLine(line)
		if err != nil {
			return nil, err
		}

		reports = append(reports, report)
	}

	return reports, nil
}

func ScanReportLine(line string) (Report, error) {
	split := strings.Split(line, " ")

	report := make(Report, 0)

	for _, s := range split {
		v, err := strconv.Atoi(s)
		if err != nil {
			return report, err
		}

		report = append(report, v)
	}

	return report, nil
}

func CountSafeReports(reports []Report) int {
	count := 0
	for _, report := range reports {
		if report.IsSafe() {
			count++
		}
	}

	return count
}

func (report *Report) IsSafe() bool {
	if len(*report) <= 1 {
		return false
	}

	initDir := GetDirection((*report)[0], (*report)[1])
	if initDir == DirectionUnknown {
		return false
	}

	previous := (*report)[0]
	for i := 1; i < len(*report); i++ {
		current := (*report)[i]

		currDirection := GetDirection(previous, current)
		if currDirection != initDir {
			return false
		}

		deviation := int(math.Abs(float64(current - previous)))
		if deviation == 0 || deviation > 3 {
			return false
		}

		previous = current
	}

	return true
}

func GetDirection(x, y int) Direction {
	if x > y {
		return DirectionDown
	}
	if x < y {
		return DirectionUp
	}
	return DirectionUnknown
}
