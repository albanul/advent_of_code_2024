package day17_test

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/albanul/advent_of_code_2024/day17"
)

func TestGame_gives_correct_output(t *testing.T) {
	table := []struct {
		regA, regB, regC int
		input            []int
		expected         string
	}{
		{10, 0, 0, []int{5, 0, 5, 1, 5, 4}, "0,1,2"},
		{2024, 0, 0, []int{0, 1, 5, 4, 3, 0}, "4,2,5,6,7,7,7,7,3,1,0"},
		{729, 0, 0, []int{0, 1, 5, 4, 3, 0}, "4,6,3,5,6,3,5,2,1,0"},
	}

	for i, scenario := range table {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			game := day17.NewGame(scenario.regA, scenario.regB, scenario.regC, scenario.input)
			expected := scenario.expected

			err := game.Play()
			if err != nil {
				t.Error("Failed to play game")
			}

			actual := game.GetOutput()
			if actual != expected {
				t.Errorf("Expected %q, got %q", expected, actual)
			}
		})
	}
}

func TestGame_has_correct_register_values(t *testing.T) {
	/*
		If register C contains 9, the program 2,6 would set register B to 1.
		If register A contains 2024, the program 0,1,5,4,3,0 would output 4,2,5,6,7,7,7,7,3,1,0 and leave 0 in register A.
		If register B contains 29, the program 1,7 would set register B to 26.
		If register B contains 2024 and register C contains 43690, the program 4,0 would set register B to 44354.
	*/

	table := []struct {
		regA, regB, regC int
		input            []int
		expected         []int
	}{
		{0, 0, 9, []int{2, 6}, []int{0, 1, 9}},
		{2024, 0, 0, []int{0, 1, 5, 4, 3, 0}, []int{0, 0, 0}},
		{0, 29, 0, []int{1, 7}, []int{0, 26, 0}},
		{0, 2024, 43690, []int{4, 0}, []int{0, 44354, 43690}},
	}

	for i, scenario := range table {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			game := day17.NewGame(scenario.regA, scenario.regB, scenario.regC, scenario.input)
			expected := scenario.expected

			err := game.Play()
			if err != nil {
				t.Error("Failed to play game")
			}

			actual := game.DropRegisters()

			if !reflect.DeepEqual(actual, expected) {
				t.Errorf("Expected %v, got %v", expected, actual)
			}
		})
	}
}
