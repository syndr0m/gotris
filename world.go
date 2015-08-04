package main

import (
	"math/rand"
	//"time"
)

const GRID_WIDTH uint = 10
const GRID_HEIGHT uint = 17

type World struct {
	score         uint
	level         uint
	lines         uint
	currentPiece  *Piece
	currentPieceX int
	currentPieceY int
	nextPiece     *Piece
	grid          [] /*Y:17*/ [] /*X:10*/ *Block
	onDeleted     func(lines uint)
}

// initialisation
func (world *World) InitGrid() {
	world.grid = make([][]*Block, GRID_HEIGHT)
	for y := uint(0); y < GRID_HEIGHT; y++ {
		world.grid[y] = make([]*Block, GRID_WIDTH)
	}
}

func (world *World) PickupPiece() *Piece {
	pieces := []*Piece{
		&PieceBar,
		&PieceCube, &PieceCube,
		&PieceL1, &PieceL1,
		&PieceL2, &PieceL2,
		&PieceTri, &PieceTri,
	}
	return pieces[rand.Intn(len(pieces))]
}
func (world *World) NextPiece() {
	world.currentPiece = world.PickupPiece()
	world.currentPieceX = 5
	world.currentPieceY = 0
}
func (world *World) Collide(piece *Piece, x int, y int) bool {
	pieceBlocks := piece.GetBlocks()
	pieceSize := piece.GetSize()
	grid := world.GetGrid()
	for tmpy := uint(0); tmpy < pieceSize; tmpy++ {
		for tmpx := uint(0); tmpx < pieceSize; tmpx++ {
			if pieceBlocks[tmpy][tmpx] != nil {
				// collision avec un bord
				if int(tmpx)+x < 0 ||
					int(tmpx)+x > 9 ||
					int(tmpy)+y > 16 {
					return true
				}
				// collision avec un block deja present
				if grid[y+int(tmpy)][x+int(tmpx)] != nil {
					return true
				}
			}
		}
	}
	return false
}
func (world *World) AttachPieceToGrid(piece *Piece, x int, y int) {
	// integrating piece inside grid
	pieceBlocks := piece.GetBlocks()
	pieceSize := piece.GetSize()
	for tmpy := uint(0); tmpy < pieceSize; tmpy++ {
		for tmpx := uint(0); tmpx < pieceSize; tmpx++ {
			if pieceBlocks[tmpy][tmpx] != nil {
				world.grid[y+int(tmpy)][x+int(tmpx)] = pieceBlocks[tmpy][tmpx]
			}
		}
	}
}

func (world *World) DeleteLines() uint {
	gridWidth := world.GetGridWidth()
	gridHeight := world.GetGridHeight()
	lines := uint(0)
	// do not check first line (avoid uint bug)
	for y := gridHeight - 1; y > 0; {
		every := true
		for x := uint(0); x < gridWidth && every; x++ {
			if world.grid[y][x] == nil {
				every = false
			}
		}
		if every {
			// we have found a line, pushing down the whole grid.
			lines++
			for y2 := y; y2 > 0; y2-- {
				for x := uint(0); x < gridWidth; x++ {
					world.grid[y2][x] = world.grid[y2-1][x]
				}
			}
			// first line is empty
			for x := uint(0); x < gridWidth; x++ {
				world.grid[0][x] = nil
			}
		} else {
			y--
		}
	}
	// Event
	if world.onDeleted != nil && lines > 0 {
		world.onDeleted(lines)
	}
	//
	return lines
}

// internals
func (world *World) CanMoveDown() bool {
	return !world.Collide(world.currentPiece, world.currentPieceX, world.currentPieceY+1)
}
func (world *World) CanMoveRight() bool {
	return !world.Collide(world.currentPiece, world.currentPieceX+1, world.currentPieceY)
}
func (world *World) CanMoveLeft() bool {
	return !world.Collide(world.currentPiece, world.currentPieceX-1, world.currentPieceY)
}
func (world *World) CanRotateRight() bool {
	world.currentPiece.RotateRight()
	collide := world.Collide(world.currentPiece, world.currentPieceX, world.currentPieceY)
	world.currentPiece.RotateLeft()
	return !collide
}
func (world *World) CanRotateLeft() bool {
	world.currentPiece.RotateLeft()
	collide := world.Collide(world.currentPiece, world.currentPieceX, world.currentPieceY)
	world.currentPiece.RotateRight()
	return !collide
}

// user controls
func (world *World) Down() {
	if world.CanMoveDown() {
		world.currentPieceY++
	} else {
		world.AttachPieceToGrid(world.currentPiece, world.currentPieceX, world.currentPieceY)
		world.DeleteLines()
		world.NextPiece()
	}
}
func (world *World) Right() {
	if world.CanMoveRight() {
		world.currentPieceX++
	}
}
func (world *World) Left() {
	if world.CanMoveLeft() {
		world.currentPieceX--
	}
}
func (world *World) RotateRight() {
	if world.CanRotateRight() {
		world.currentPiece.RotateRight()
	}
}
func (world *World) RotateLeft() {
	if world.CanRotateLeft() {
		world.currentPiece.RotateLeft()
	}
}
func (world *World) Space() { /* FIXME */ }

func (world *World) GetGrid() [][]*Block { return world.grid }
func (world *World) GetGridWidth() uint  { return GRID_WIDTH }
func (world *World) GetGridHeight() uint { return GRID_HEIGHT }

func (world *World) GetPiece() *Piece { return world.currentPiece }
func (world *World) GetPieceX() int   { return world.currentPieceX }
func (world *World) GetPieceY() int   { return world.currentPieceY }

// Events
func (world *World) OnDeleted(f func(uint)) { world.onDeleted = f }

// Creating the world
func NewWorld() *World {
	world := new(World)
	world.InitGrid()
	world.level = 1
	world.lines = 0
	world.score = 0
	world.NextPiece()

	// every seconds, we drop a piece
	/*ticker := time.NewTicker(time.Millisecond * time.Duration(1000))
	go func() {
		for {
			<-ticker.C
			// down.
			world.Down()
		}
	}()*/

	return world
}