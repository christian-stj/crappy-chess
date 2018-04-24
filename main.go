package main

import (
"github.com/hajimehoshi/ebiten"
	"image/color"
)


func update(screen *ebiten.Image) error {
	var board [8][8]int
	i := 30
	img,_ := ebiten.NewImage(i,i, 0)


	for a, row := range(board) {
		for b, _ := range(row) {
			op := &ebiten.DrawImageOptions{}
			if (b%2 == 0 && a%2 == 0 || b%2 != 0 && a%2 != 0) {
				img.Fill(color.White)
			} else {
				img.Fill(color.NRGBA{100,250,100,100})
			}
			op.GeoM.Translate(float64(10+i+i*(b-1)), float64(10+i+i*(a-1)))
			screen.DrawImage(img, op)
		}
	}

	return nil
}

func main() {
	ebiten.Run(update, 260, 260, 2, "Chess")
}