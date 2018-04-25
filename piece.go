package main

import (
	"github.com/hajimehoshi/ebiten"
	_ "image/png"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Piece struct{
	pieceIcon *ebiten.Image
	color int
	rank int

}

func CreatePiece(color int, rank int)(p Piece){
	p.color = color
	p.rank = rank
	img, _, _ := ebitenutil.NewImageFromFile("icons/icons8-bishop-2.png", 0)
	p.pieceIcon = img

	return p
}