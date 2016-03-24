package mandelbrot

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"io"
	"fmt"
)

func Mandelbrot(out io.Writer, magnification float64) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
	)
	var width, height int
	width = 1024
	height = 1024

	if magnification == 0 {
		// 指定されなかったとみなし何もしない
	}else if magnification < 0.1 || magnification > 3 {
		fmt.Println("unsupported magnification value")
		return
	}else {
		width = int(float64(width) * magnification)
		height = int(float64(height) * magnification)
	}


	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			// 画像の点(px、py)は複素数zを表している。
			img.Set(px, py, mandelbrot(z))
		}
	}

	png.Encode(out, img) // 注意： エラーを無視
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{255 - contrast*n,0,0,255}
		}
	}
	return color.Black
}
