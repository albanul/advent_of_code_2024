package day17

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type fn func(game *Game, op int)

type Game struct {
	instructions     map[int]fn
	index            int
	regA, regB, regC int
	input            []int
	output           []int
}

func (g *Game) Play() error {
	for g.index < len(g.input) {
		opCode := g.input[g.index]
		op := g.input[g.index+1]

		f, ok := g.instructions[opCode]
		if !ok {
			return fmt.Errorf("no instruction for opcode %d", opCode)
		}

		f(g, op)
	}

	return nil
}

func (g *Game) DropRegisters() []int {
	return []int{g.regA, g.regB, g.regC}
}

func NewGame(regA, regB, regC int, input []int) *Game {
	instructions := map[int]fn{
		0: adv,
		1: bxl,
		2: bst,
		3: jnz,
		4: bxc,
		5: out,
		6: bdv,
		7: cdv,
	}

	return &Game{
		index:        0,
		regA:         regA,
		regB:         regB,
		regC:         regC,
		input:        input,
		instructions: instructions,
		output:       make([]int, 0),
	}
}

func (g *Game) literalOp(v int) int {
	return v
}

func (g *Game) comboOp(v int) int {
	if v >= 0 && v <= 3 {
		return v
	}

	if v == 4 {
		return g.regA
	}

	if v == 5 {
		return g.regB
	}

	if v == 6 {
		return g.regC
	}

	panic("invalid combo op")
}

func (g *Game) GetOutput() string {
	arr := make([]string, 0)
	for _, v := range g.output {
		arr = append(arr, strconv.Itoa(v))
	}
	return strings.Join(arr, ",")
}

func adv(game *Game, op int) {
	numerator := float64(game.regA)

	power := game.comboOp(op)
	denominator := math.Pow(2, float64(power))

	result := int(math.Floor(numerator / denominator))

	game.regA = result

	game.index += 2
}

func bxl(game *Game, op int) {
	l := game.regB
	r := game.literalOp(op)

	game.regB = l ^ r

	game.index += 2
}

func bst(game *Game, op int) {
	game.regB = game.comboOp(op) % 8

	game.index += 2
}

func jnz(game *Game, op int) {
	if game.regA == 0 {
		game.index += 2
		return
	}

	game.index = game.literalOp(op)
}

func bxc(game *Game, _ int) {
	game.regB = game.regB ^ game.regC
	game.index += 2
}

func out(game *Game, op int) {
	v := game.comboOp(op) % 8
	game.output = append(game.output, v)
	game.index += 2
}

func bdv(game *Game, op int) {
	numerator := float64(game.regA)

	power := game.comboOp(op)
	denominator := math.Pow(2, float64(power))

	result := int(math.Ceil(numerator / denominator))

	game.regB = result

	game.index += 2
}

func cdv(game *Game, op int) {
	numerator := float64(game.regA)

	power := game.comboOp(op)
	denominator := math.Pow(2, float64(power))

	result := int(math.Ceil(numerator / denominator))

	game.regC = result

	game.index += 2
}
