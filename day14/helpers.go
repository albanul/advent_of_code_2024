package day14

type Point struct {
	I, J int
}

type Velocity struct {
	Di, Dj int
}

type Robot struct {
	Position Point
	Velocity Velocity
}

func (r *Robot) move(width, height int) {
	newI := ((r.Position.I+r.Velocity.Di)%height + height) % height
	newJ := ((r.Position.J+r.Velocity.Dj)%width + width) % width

	r.Position.I = newI
	r.Position.J = newJ
}

func newRobot(point Point, velocity Velocity) Robot {
	return Robot{point, velocity}
}

type Game struct {
	robots        []*Robot
	width, height int
}

func (g *Game) GetRobots() []Robot {
	result := make([]Robot, 0)

	for _, r := range g.robots {
		nr := newRobot(r.Position, r.Velocity)
		result = append(result, nr)
	}

	return result
}

func (g *Game) PlayFor(seconds int) {
	for range seconds {
		g.PlayFor1Second()
	}
}

func (g *Game) PlayFor1Second() {
	for _, r := range g.robots {
		r.move(g.width, g.height)
	}
}

func (g *Game) GetWidth() int {
	return g.width
}

func (g *Game) GetHeight() int {
	return g.height
}

func NewGame(width, height int, robots []*Robot) *Game {
	return &Game{robots, width, height}
}

func CalculateSafetyFactor(game *Game) int {
	hw := game.width / 2
	hh := game.height / 2

	tlc, trc, brc, blc := 0, 0, 0, 0

	for _, robot := range game.robots {
		ri, rj := robot.Position.I, robot.Position.J
		if ri < hh && rj < hw {
			tlc++
		}

		if ri < hh && rj > hw {
			trc++
		}

		if ri > hh && rj > hw {
			brc++
		}

		if ri > hh && rj < hw {
			blc++
		}
	}

	return tlc * trc * brc * blc
}
