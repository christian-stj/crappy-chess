package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"image/color"
)

var board *[8][8]Tile
var selectedTile *Tile
var previousColor color.Color
var playersTurn = 0
var check = false

// updates the screen
func update(screen *ebiten.Image) error {
	if selectedTile == nil {
		selectPiece()
	} else {
		movePiece()
	}
	UpdateBoard(screen, board)

	return nil
}

// reads click input
func click() (bool, int, int) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		return true, x, y
	} else {
		return false, 0, 0
	}
}

// Main routine, runs game
func main() {
	board = CreateBoard()
	StartingPiecePos(board)
	fmt.Println(board[3][3].piece.color) //This is the problem
	ebiten.Run(update, 280, 280, 2, "Chess")
}
