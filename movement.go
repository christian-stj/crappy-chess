package main

import "image/color"

func selectTile() {
	if flag, x, y:=click(); flag == true {
		for a, row := range(board) {
			for b,_ := range(row) {
				tile:=&board[a][b]
				if (tile.xleft <= x && x < tile.xright) && (tile.ytop <= y && y < tile.ybot) && tile.piece!=(Piece{}) && tile.piece.color==playersTurn {
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
}

func moveTile(){
	if flag, x, y:=click(); flag == true {
		for a, row := range(board) {
			for b,_ := range(row) {
				tile:=&board[a][b]
				if (tile.xleft <= x && x < tile.xright) && (tile.ytop <= y && y < tile.ybot) {
					if selectedTile == tile {
						selectedTile.color = previousColor
						selectedTile = nil
					} else {
						if canMove(tile) {
							tile.piece = selectedTile.piece
							selectedTile.piece = Piece{}
							selectedTile.color = previousColor
							selectedTile = nil
							changeTurn()
						} else {

						}
					}
				}
			}
		}
	}

}

func changeTurn(){
	if playersTurn == 0{
		playersTurn = 1
	} else {
		playersTurn = 0
	}
}

func canMove(tile *Tile) bool{
	var boo bool
	switch selectedTile.piece.rank {
	case 0:
	case 1:
		boo = movePeasant(tile)
	case 2:
	case 3:
	case 4:
	case 5:

	}
	return boo
}

func movePeasant(tile *Tile) bool {
	return false
}
