package main

import "image/color"

func selectTile() {
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
						tile.piece = selectedTile.piece
						selectedTile.piece = Piece{}
						selectedTile.color = previousColor
						selectedTile = nil
					}
				}
			}
		}
	}

}
