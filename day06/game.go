package day06

type Direction int

const (
	Top Direction = iota
	Right
	Bottom
	Left
)

type Position struct {
	X, Y int
}

type Guard struct {
	Position  Position
	Direction Direction
}

type Game struct {
	width, height   int
	counter         int
	patrolPath      map[Position]bool
	walls           map[Position]bool
	initialGuard    Guard
	guard           Guard
	isInLoop        bool
	lastNewPosition Position
}

func (g *Game) CanSetWall(x, y int) bool {
	position := Position{x, y}
	_, hasWall := g.walls[position]

	canSetWall := !hasWall && position != g.guard.Position
	return canSetWall
}

func (g *Game) SetWall(x, y int) {
	position := Position{x, y}
	g.walls[position] = true
}

func (g *Game) SetWalls(walls map[Position]bool) {
	g.walls = walls
}

func (g *Game) SetGuard(guard Guard) {
	g.counter = 1
	g.initialGuard = guard
	g.guard = guard
	g.patrolPath[g.guard.Position] = true
	g.lastNewPosition = g.guard.Position
}

func (g *Game) MakeMove() {
	deltaX, deltaY := 0, 0
	if g.guard.Direction == Top {
		deltaX, deltaY = 0, -1
	}

	if g.guard.Direction == Bottom {
		deltaX, deltaY = 0, 1
	}

	if g.guard.Direction == Left {
		deltaX, deltaY = -1, 0
	}

	if g.guard.Direction == Right {
		deltaX, deltaY = 1, 0
	}

	newPosition := Position{g.guard.Position.X + deltaX, g.guard.Position.Y + deltaY}
	isWall := g.walls[newPosition]

	if isWall {
		newDirection := (g.guard.Direction + 1) % 4
		g.guard.Direction = newDirection
	} else {
		wasVisited := g.patrolPath[newPosition]
		if !wasVisited {
			g.patrolPath[newPosition] = true
			g.counter++
			g.lastNewPosition = newPosition
		} else {
			if newPosition == g.lastNewPosition {
				g.isInLoop = true
			}
		}
		g.guard.Position = newPosition
	}
}

func (g *Game) GetCounter() int {
	return g.counter - 1
}

func (g *Game) IsOver() bool {
	return g.guard.Position.X >= g.width || g.guard.Position.X < 0 ||
		g.guard.Position.Y >= g.height || g.guard.Position.Y < 0
}

func (g *Game) IsInLoop() bool {
	return g.isInLoop
}

func (g *Game) GetSizes() (width int, height int) {
	return g.width, g.height
}

func (g *Game) Reset(addedWallPosition Position) {
	delete(g.walls, addedWallPosition)

	g.guard = g.initialGuard
	g.patrolPath = make(map[Position]bool)
	g.patrolPath[g.guard.Position] = true
	g.counter = 1
	g.lastNewPosition = g.guard.Position
	g.isInLoop = false
}

func NewGame(width, height int) *Game {
	game := &Game{
		width:      width,
		height:     height,
		walls:      make(map[Position]bool),
		patrolPath: make(map[Position]bool),
	}

	return game
}
