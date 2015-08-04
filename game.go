package main

import (
	"fmt"
	"mygameengine"
	"mygameengine/image"
)

const IMG_BLOCK_01 string = "assets/images/block-01.png"
const IMG_BLOCK_02 string = "assets/images/block-02.png"
const IMG_BLOCK_03 string = "assets/images/block-03.png"
const IMG_BLOCK_04 string = "assets/images/block-04.png"
const IMG_BLOCK_05 string = "assets/images/block-05.png"
const IMG_BLOCK_06 string = "assets/images/block-06.png"

const BLOCK_SIZE uint = 27

const KEY_RIGHT int = 57
const KEY_LEFT int = 58
const KEY_UP int = 60
const KEY_DOWN int = 59
const KEY_CTRL_RIGHT int = 104
const KEY_CTRL_LEFT int = 100

func keyHandler(world *World) func(int) {
	return func(key int) {
		switch key {
		case KEY_RIGHT:
			fmt.Println("move right")
			world.CanMoveRight()
			world.Right()
		case KEY_LEFT:
			fmt.Println("move left")
			world.CanMoveLeft()
			world.Left()
		case KEY_DOWN:
			fmt.Println("move down")
			world.CanMoveDown()
			world.Down()
		case KEY_CTRL_RIGHT, KEY_UP:
			fmt.Println("rotate right")
			world.CanRotateRight()
			world.RotateRight()
		case KEY_CTRL_LEFT:
			world.CanRotateLeft()
			world.RotateLeft()
		}
		fmt.Println("game: KEY DOWN", key)
	}
}

func drawGameBackground(screen *image.Image) {
	screen.DrawRectangle(0, 0, 640, 480, mygameengine.COLOR_BLACK)
}

func drawGameGrid(screen *image.Image, world *World, engine *mygameengine.MyGameEngine) {
	gridWidth := world.GetGridWidth()
	gridHeight := world.GetGridHeight()
	blocks := world.GetGrid()
	for y := uint(0); y < gridHeight; y++ {
		for x := uint(0); x < gridWidth; x++ {
			if blocks[y][x] != nil {
				image := blockToImage(blocks[y][x], engine)
				screen.BlitAt(image, int(BLOCK_SIZE+x*BLOCK_SIZE), int(y*BLOCK_SIZE))
			}
		}
	}
}

func drawGameGridBorders(screen *image.Image, world *World, engine *mygameengine.MyGameEngine) {
	gridWidth := world.GetGridWidth()
	gridHeight := world.GetGridHeight()
	for y := uint(0); y < gridHeight; y++ {
		screen.BlitAt(engine.Assets().Get(IMG_BLOCK_05), 0, int(y*BLOCK_SIZE))
		screen.BlitAt(engine.Assets().Get(IMG_BLOCK_06), int(gridWidth*BLOCK_SIZE+BLOCK_SIZE), int(y*BLOCK_SIZE))
	}
	for x := uint(0); x < gridWidth+2; x++ {
		screen.BlitAt(engine.Assets().Get(IMG_BLOCK_06), int(x*BLOCK_SIZE), int(gridHeight*BLOCK_SIZE))
	}
}

func drawGameCurrentPiece(screen *image.Image, world *World, engine *mygameengine.MyGameEngine) {
	piece := world.GetPiece()
	pieceBlocks := piece.GetBlocks()
	pieceSize := int(piece.GetSize())
	pieceY := world.GetPieceY()
	pieceX := world.GetPieceX()

	for y := 0; y < pieceSize; y++ {
		for x := 0; x < pieceSize; x++ {
			if pieceBlocks[y][x] != nil {
				image := blockToImage(pieceBlocks[y][x], engine)
				screen.BlitAt(
					image,
					int(BLOCK_SIZE+uint(pieceX+x)*BLOCK_SIZE),
					int(uint(pieceY+y)*BLOCK_SIZE),
				)
			}
		}
	}
}

func drawGameNextPiece(screen *image.Image, world *World, engine *mygameengine.MyGameEngine) {
	piece := world.GetNextPiece()
	pieceBlocks := piece.GetBlocks()
	pieceSize := piece.GetSize()

	marginLeft := uint(15)
	marginTop := uint(8)
	width := uint(6)

	// border
	for y := uint(0); y < width; y++ {
		for x := uint(0); x < width; x++ {
			if x == 0 || y == 0 || x == 5 || y == 5 {
				screen.BlitAt(
					engine.Assets().Get(IMG_BLOCK_05),
					int(BLOCK_SIZE*marginLeft+uint(x)*BLOCK_SIZE),
					int(BLOCK_SIZE*marginTop+uint(y)*BLOCK_SIZE),
				)
			}
		}
	}
	// piece
	for y := uint(0); y < pieceSize; y++ {
		for x := uint(0); x < pieceSize; x++ {
			if pieceBlocks[y][x] != nil {
				image := blockToImage(pieceBlocks[y][x], engine)
				screen.BlitAt(
					image,
					int(BLOCK_SIZE*(marginLeft+1)+uint(x)*BLOCK_SIZE),
					int(BLOCK_SIZE*(marginTop+1)+uint(y)*BLOCK_SIZE),
				)
			}
		}
	}
}

func drawGameFlash(screen *image.Image, flash int) {
	white := mygameengine.COLOR_WHITE
	white.A = uint8(25 * flash)
	screen.DrawRectangle(0, 0, 640, 480, white)
}

func repaintHandler(engine *mygameengine.MyGameEngine, world *World) func(*image.Image) {
	// flash effect setup
	flash := 0
	world.OnDeleted(func(lines uint) {
		flash = 10
	})

	return func(screen *image.Image) {
		drawGameBackground(screen)
		drawGameGrid(screen, world, engine)
		drawGameGridBorders(screen, world, engine)
		drawGameCurrentPiece(screen, world, engine)
		drawGameNextPiece(screen, world, engine)
		// flash effect on line deletion
		if flash > 0 {
			drawGameFlash(screen, flash)
			flash--
		}
	}
}

func loadAssets(engine *mygameengine.MyGameEngine) {
	engine.Assets().Png(IMG_BLOCK_01)
	engine.Assets().Png(IMG_BLOCK_02)
	engine.Assets().Png(IMG_BLOCK_03)
	engine.Assets().Png(IMG_BLOCK_04)
	engine.Assets().Png(IMG_BLOCK_05)
	engine.Assets().Png(IMG_BLOCK_06)
}

func blockToImage(block *Block, engine *mygameengine.MyGameEngine) *image.Image {
	switch block.GetShape() {
	case 1:
		return engine.Assets().Get(IMG_BLOCK_01)
	case 2:
		return engine.Assets().Get(IMG_BLOCK_02)
	case 3:
		return engine.Assets().Get(IMG_BLOCK_03)
	case 4:
		return engine.Assets().Get(IMG_BLOCK_04)
	}
	return nil
}

func Game(engine *mygameengine.MyGameEngine) *mygameengine.Board {
	loadAssets(engine)
	gameBoard := mygameengine.NewBoard()
	world := NewWorld()
	gameBoard.OnStart(world.Start)
	gameBoard.OnKeyDown(keyHandler(world))
	gameBoard.OnRepaint(repaintHandler(engine, world))
	return gameBoard
}
