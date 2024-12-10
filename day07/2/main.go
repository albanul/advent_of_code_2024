package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	Target    int64
	Numbers   []int
	Operators int
}

func (e *Equation) Calculate() int64 {
	res := int64(e.Numbers[0])

	op := e.Operators
	power := 1

	end := len(e.Numbers) - 1

	for power <= end {
		check := op % 3

		right := int64(e.Numbers[power])

		if check == 0 {
			res += right
		}

		if check == 1 {
			res *= right
		}

		if check == 2 {
			powerOfTen := 1
			for int64(powerOfTen) <= right {
				powerOfTen *= 10
			}

			res = res*int64(powerOfTen) + right
		}

		op /= 3
		power++
	}

	return res
}

func (e *Equation) IsCorrect() bool {
	return e.Calculate() == e.Target
}

func (e *Equation) SetOperators(operators int) {
	e.Operators = operators
}

type Puzzle struct {
	Equations []Equation
}

func main() {
	puzzle, err := GetPuzzleFromFile("day07/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum := GetSumOfSolvableEquations(puzzle)
	fmt.Println("Sum: ", sum)
}

func GetPuzzleFromFile(filename string) (*Puzzle, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	equations := make([]Equation, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, ":")

		res, err := strconv.ParseInt(split[0], 10, 64)
		if err != nil {
			return nil, errors.Join(errors.New("can't parse target result"), err)
		}

		numbers := make([]int, 0)

		split[1] = strings.TrimLeft(split[1], " ")
		rawNumbers := strings.Split(split[1], " ")
		for _, rawNumber := range rawNumbers {
			number, err := strconv.Atoi(rawNumber)
			if err != nil {
				return nil, errors.Join(errors.New("can't parse a number"), err)
			}

			numbers = append(numbers, number)
		}

		equation := Equation{res, numbers, 0}
		equations = append(equations, equation)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	puzzle := &Puzzle{equations}

	return puzzle, nil
}

func GetSumOfSolvableEquations(puzzle *Puzzle) int64 {
	sum := int64(0)
	for _, equation := range puzzle.Equations {
		isSolvable := CanSolve(equation)

		if isSolvable {
			sum += equation.Target
		}
	}

	return sum
}

func CanSolve(equation Equation) bool {
	if equation.IsCorrect() {
		return true
	}

	end := int(math.Pow(3, float64(len(equation.Numbers)-1)))

	for i := 1; i < end; i++ {
		equation.SetOperators(i)

		if equation.IsCorrect() {
			return true
		}
	}

	return false
}
