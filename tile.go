package main

import (
	"image/color"
)

type Tile struct {
	xleft  int
	xright int
	ytop   int
	ybot   int
	a      int
	b      int
	color  color.Color
	piece  Piece
}

func CreateTile(xleft int, xright int, ytop int, ybot int) (t Tile) {
	tile := Tile{}
	tile.xleft = xleft
	tile.xright = xright
	tile.ytop = ytop
	tile.ybot = ybot
	return tile
}
