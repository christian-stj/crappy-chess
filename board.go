package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"image"
)

type Tile struct {
	xleft int
	xright int
	ytop int
	ybot int
	color color.Color
	piece image.Image
}

func CreateBoard(screen *ebiten.Image) ([8][8]Tile){
	var board [8][8]Tile
	i := 30
	img,_ := ebiten.NewImage(i,i, 0)


	for a, row := range(board) {
		for b, tile := range(row) {
			op := &ebiten.DrawImageOptions{}
			if (b%2 == 0 && a%2 == 0 || b%2 != 0 && a%2 != 0) {
				tile.color=color.White
			} else {
				tile.color=color.NRGBA{100,250,100,100}
			}
			tile.xleft=10+i+i*(b-1)
			tile.xright=tile.xleft+i
			tile.ytop=10+i+i*(a-1)
			tile.ybot=tile.ytop+i
			img.Fill(tile.color)
			op.GeoM.Translate(float64(10+i+i*(b-1)), float64(10+i+i*(a-1)))
			screen.DrawImage(img, op)
		}
	}
	return board
}