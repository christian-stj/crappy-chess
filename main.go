package main

import (
"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

var board *[8][8]Tile


func update(screen *ebiten.Image) error {

	UpdateBoard(screen, board)
	StartingPiecePos(board)
	click()
	return nil
}

func click() bool {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
	}
	return false
}

func main() {
	board = CreateBoard()
	ebiten.Run(update, 260, 260, 2, "Chess")

}