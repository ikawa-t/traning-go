package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// 画像の点(px、py)は複素数zを表している。
			img.Set(px, py, kon(z))
		}
	}

	//標準出力だとうまくいかない
	//png.Encode(os.Stdout, img) // 注意： エラーを無視
	//ファイル出力
	f, _ := os.OpenFile("data/kon.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)

}

func kon(z complex128) color.Color {
	const iterations = 200
	const contrast = 10
	const eps =0.01             //収束条件

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		// ニュートン法
		v = v - ((z*z*z*z-1) / (4*z*z*z))
		// 絶対値が解の絶対値1に近づいているかの判定
		if(math.Abs(cmplx.Abs(v) - 1) < eps) {
			return color.RGBA{255 - contrast*n,0,0,255}
		}
	}
	return color.Black
}
