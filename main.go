package main

import (
"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"fmt"
	"image/color"
)

var board *[8][8]Tile
var selectedTile *Tile
var previousColor color.Color


func update(screen *ebiten.Image) error {

	UpdateBoard(screen, board)
	StartingPiecePos(board)
	if flag, x, y:=click(); flag == true {
		for a, row := range(board) {
			for b,_ := range(row) {
				tile:=&board[a][b]
				if (tile.xleft <= x && x < tile.xright) && (tile.ytop <= y && y < tile.ybot) && tile.piece!=(Piece{}) {
					if selectedTile != nil {
						selectedTile.color=previousColor
					}
					previousColor=tile.color
					tile.color=color.RGBA{250, 100, 50, 200}
					selectedTile=tile
				}
			}
		}
	}
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
	ebiten.Run(update, 260, 260, 2, "Chess")

}