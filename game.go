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

func keyHandler(world *World) func(int) {
	return func(key int) {
		switch key {
		case 57:
			fmt.Println("move right")
			world.CanMoveRight()
			world.Right()
		case 58:
			fmt.Println("move left")
			world.CanMoveLeft()
			world.Left()
		case 59:
			fmt.Println("move down")
			world.CanMoveDown()
			world.Down()
		case 104:
			fmt.Println("rotate right")
			world.CanRotateRight()
			world.RotateRight()
			/*case 58:
			world.CanRotateLeft()
			world.RotateLeft()*/
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

func repaintHandler(engine *mygameengine.MyGameEngine, world *World) func(*image.Image) {
	return func(screen *image.Image) {
		drawGameBackground(screen)
		drawGameGrid(screen, world, engine)
		drawGameGridBorders(screen, world, engine)
		drawGameCurrentPiece(screen, world, engine)
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
	world.OnDeleted(func(lines uint) { fmt.Println("deleted ", lines) })
	gameBoard.OnKeyDown(keyHandler(world))
	gameBoard.OnRepaint(repaintHandler(engine, world))
	return gameBoard
}
