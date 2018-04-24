package main

import (
"github.com/hajimehoshi/ebiten"
)


func update(screen *ebiten.Image) error {

	CreateBoard(screen)

	return nil
}

func main() {
	ebiten.Run(update, 260, 260, 2, "Chess")
}