package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	// 画像の読み込み
	file, err := os.Open("data/mandelbrot.png")
	//defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(err)
		fmt.Println("error")
		return
	}

	fmt.Println("hogehoge")


	// 変換
	rect := img.Bounds()
	width := rect.Size().X
	height := rect.Size().Y
	rgba := image.NewRGBA(rect)

	for x := 1; x < width-1; x++ {
		for y := 1; y < height-1; y++ {
			r0, g0, b0, a0 := img.At(x, y).RGBA()
			r1, g1, b1, a1 := img.At(x-1, y).RGBA()
			r2, g2, b2, a2 := img.At(x+1, y).RGBA()
			r3, g3, b3, a3 := img.At(x, y-1).RGBA()
			r4, g4, b4, a4 := img.At(x, y+1).RGBA()

			var col color.RGBA
			col.R = uint8((r0+r1+r2+r3+r4)/5)
			col.G = uint8((g0+g1+g2+g3+g4)/5)
			col.B = uint8((b0+b1+b2+b3+b4)/5)
			col.A = uint8((a0+a1+a2+a3+a4)/5)
			rgba.Set(x, y, col)
		}
	}

	// 画像の書き込み
	f, _ := os.OpenFile("data/sampling.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, rgba)


}
