package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
)


func CreateBoard()(*[8][8]Tile){
	var board [8][8]Tile
	i := 30
	for a, row := range(board) {
		for b, _ := range (row) {
			xleft := 10 + i + i*(b-1)
			xright := xleft + i
			ytop := 10 + i + i*(a-1)
			ybot := ytop + i
			tile := CreateTile(xleft,xright,ytop,ybot)
			if b%2 == 0 && a%2 == 0 || b%2 != 0 && a%2 != 0 {
				tile.color = color.White
			} else {
				tile.color = color.NRGBA{100, 250, 100, 100}
			}
			board[a][b] = tile
		}
	}
	return &board
}

func UpdateBoard(screen *ebiten.Image, board *[8][8]Tile){
	i := 30
	img,_ := ebiten.NewImage(i,i, 0)

	for a, row := range(board) {
		for b, tile := range(row) {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(10 + i + i*(b-1)), float64(10 + i + i*(a-1)))
			img.Fill(tile.color)
			screen.DrawImage(img, op)
			if tile.piece != (Piece{}) {
				screen.DrawImage(tile.piece.pieceIcon, op)
			}
		}
	}
}

func StartingPiecePos(board *[8][8]Tile){
	board[0][0].piece = CreatePiece(1,3)

}