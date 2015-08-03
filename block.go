package main

type Block struct {
	shape uint
}

func (block *Block) GetShape() uint { return block.shape }

var Block01 Block = Block{shape: 1}
var Block02 Block = Block{shape: 2}
var Block03 Block = Block{shape: 3}
var Block04 Block = Block{shape: 4}
var Block05 Block = Block{shape: 5}
var Block06 Block = Block{shape: 6}
