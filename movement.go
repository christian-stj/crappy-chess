package main

import "image/color"

func selectTile() {
	if flag, x, y := click(); flag == true {
		for a, row := range board {
			for b, _ := range row {
				tile := &board[a][b]
				if (tile.xleft <= x && x < tile.xright) && (tile.ytop <= y && y < tile.ybot) && tile.piece != (Piece{}) && tile.piece.color == playersTurn {
					if selectedTile != nil {
						selectedTile.color = previousColor
					}
					previousColor = tile.color
					tile.color = color.RGBA{250, 100, 50, 200}
					selectedTile = tile
				}
			}
		}
	}
}

func moveTile() {
	if flag, x, y := click(); flag == true {
		for a, row := range board {
			for b, _ := range row {
				tile := &board[a][b]
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

func changeTurn() {
	if playersTurn == 0 {
		playersTurn = 1
	} else {
		playersTurn = 0
	}
}

func canMove(tile *Tile) bool {
	var boo bool
	switch selectedTile.piece.rank {
	case 0:
		boo = movePawn(tile)
	case 1:
		boo = moveKnight(tile)
	case 2:
		boo = moveBishop(tile)
	case 3:
		boo = moveRook(tile)
	case 4:
		boo = moveQueen(tile)
	case 5:
		boo = moveKing(tile)

	}
	return boo
}

func movePawn(tile *Tile) bool {
	xold, yold := selectedTile.b, selectedTile.a
	xnew, ynew := tile.b, tile.a
	c := selectedTile.piece.color
	switch c{
	case 0:
		if xnew == xold && ynew == yold-1 && tile.piece == (Piece{}) {
			return true
		}else if xnew == xold && ynew == yold-2 && tile.piece == (Piece{}) && board[ynew+1][xold].piece == (Piece{}) && yold == 6 {
			return true
		}
	case 1:
	    if xnew == xold && ynew == yold+1 && tile.piece == (Piece{}) {
			return true
		} else if xnew == xold && ynew == yold+2 && tile.piece == (Piece{}) && board[ynew-1][xold].piece == (Piece{}) && yold == 1 {
			return true
		}

	}
	return false
}

func moveKnight(tile *Tile) bool {
	return true
}

func moveBishop(tile *Tile) bool {
	return true
}

func moveRook(tile *Tile) bool {
	return true
}

func moveQueen(tile *Tile) bool {
	return true
}

func moveKing(tile *Tile) bool {
	return true
}
