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
			tile.a=a
			tile.b=b
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
				opPiece:= &ebiten.DrawImageOptions{}
				opPiece.GeoM.Scale(0.75,0.75)
				opPiece.GeoM.Translate(float64(10 + i + i*(b-1)), float64(10 + i + i*(a-1)))
				screen.DrawImage(tile.piece.pieceIcon, opPiece)
			}
		}
	}
}

func StartingPiecePos(board *[8][8]Tile){
	// Draw the "black" pieces
	board[0][0].piece = CreatePiece(1,3)
	board[0][1].piece = CreatePiece(1,1)
	board[0][2].piece = CreatePiece(1,2)
	board[0][3].piece = CreatePiece(1,5)
	board[0][4].piece = CreatePiece(1,4)
	board[0][5].piece = CreatePiece(1,2)
	board[0][6].piece = CreatePiece(1,1)
	board[0][7].piece = CreatePiece(1,3)
	for i:=0; i<8; i++{
		board[1][i].piece = CreatePiece(1,0)
	}

	// Draw the "white" pieces
	board[7][0].piece = CreatePiece(0,3)
	board[7][1].piece = CreatePiece(0,1)
	board[7][2].piece = CreatePiece(0,2)
	board[7][3].piece = CreatePiece(0,5)
	board[7][4].piece = CreatePiece(0,4)
	board[7][5].piece = CreatePiece(0,2)
	board[7][6].piece = CreatePiece(0,1)
	board[7][7].piece = CreatePiece(0,3)
	for i:=0; i<8; i++{
		board[6][i].piece = CreatePiece(0,0)
	}

}