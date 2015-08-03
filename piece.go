package main

type Piece struct {
	blocks [] /*y*/ [] /*x*/ *Block
}

var PieceTri Piece = Piece{
	blocks: [][]*Block{
		{nil, &Block01, nil},
		{nil, &Block01, &Block01},
		{nil, &Block01, nil},
	},
}

var PieceL1 Piece = Piece{
	blocks: [][]*Block{
		{nil, &Block02, nil},
		{nil, &Block02, nil},
		{nil, &Block02, &Block02},
	},
}

var PieceL2 Piece = Piece{
	blocks: [][]*Block{
		{nil, &Block02, nil},
		{nil, &Block02, nil},
		{&Block02, &Block02, nil},
	},
}

var PieceCube Piece = Piece{
	blocks: [][]*Block{
		{&Block03, &Block03},
		{&Block03, &Block03},
	},
}

var PieceBar Piece = Piece{
	blocks: [][]*Block{
		{nil, &Block04, nil, nil},
		{nil, &Block04, nil, nil},
		{nil, &Block04, nil, nil},
		{nil, &Block04, nil, nil},
	},
}

func (piece *Piece) GetBlocks() [][]*Block {
	return piece.blocks
}
func (piece *Piece) GetSize() uint {
	return uint(len(piece.blocks))
}

/*
** Assuming piece's shapes are square.
 */
func (piece *Piece) RotateRight() {
	var result [][]*Block
	var size uint = piece.GetSize()

	result = make([][]*Block, size)
	for y := uint(0); y < size; y++ {
		result[y] = make([]*Block, size)
		for x := uint(0); x < size; x++ {
			result[y][x] = piece.blocks[size-x-1][y]
		}
	}
	piece.blocks = result
}
func (piece *Piece) RotateLeft() {
	var result [][]*Block
	var size uint = piece.GetSize()

	result = make([][]*Block, size)
	for y := uint(0); y < size; y++ {
		result[y] = make([]*Block, size)
		for x := uint(0); x < size; x++ {
			result[y][x] = piece.blocks[x][size-y-1]
		}
	}
	piece.blocks = result
}
