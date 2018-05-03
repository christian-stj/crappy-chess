package main

import (
"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"image/color"
)

var board *[8][8]Tile
var selectedTile *Tile
var previousColor color.Color
var playersTurn int = 0


func update(screen *ebiten.Image) error {
	if selectedTile == nil {
		selectTile()
	} else {
		moveTile()
	}
	UpdateBoard(screen, board)

	return nil
}


func click() (bool, int , int) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		return true, x, y
	} else {
		return false, 0, 0
	}
}

func main() {
	board = CreateBoard()
	StartingPiecePos(board)
	ebiten.Run(update, 260, 260, 2, "Chess")
}