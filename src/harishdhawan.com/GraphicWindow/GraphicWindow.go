package main

import _ "image/png"

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
)

func main() {
	fmt.Println("it works..")
	m := image.NewRGBA(image.Rect(0, 0, 640, 480))
	blue := color.RGBA{0, 0, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)
}
