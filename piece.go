package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"strconv"
)

// color : 0 = white, 1 = black
// rank : 0 = bonde, 1 = häst, 2 = löpare, 3 = torn, 4 = dam, 5 = kung
type Piece struct {
	pieceIcon *ebiten.Image
	color     int
	rank      int
}

//Creates a new Piece with the specified color and rank.
func CreatePiece(color int, rank int) (p Piece) {
	p.color = color
	p.rank = rank
	img, _, _ := ebitenutil.NewImageFromFile("icons/"+strconv.Itoa(color)+strconv.Itoa(rank)+".png", 0)
	p.pieceIcon = img

	return p
}
