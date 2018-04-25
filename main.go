package main

import (
"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)


func update(screen *ebiten.Image) error {

	CreateBoard(screen)
	click()
	return nil
}

func click() bool {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
	}
	return false
}

func main() {
	ebiten.Run(update, 260, 260, 2, "Chess")
}