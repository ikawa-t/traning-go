package lissajous

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	//"time"
)

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}}

const (
	whiteIndex = 0 // パレットの最初の色
	blackIndex = 1 // パレットの次の色
)

//func main() {
//	rand.Seed(time.Now().UTC().UnixNano())
//	lissajous(os.Stdout)
//}

func Lissajous(out io.Writer, cycles float64) {
	const (
		//cycles  = 5     // 完全なxのオシレータ回転数
		res     = 0.001 // 角度回転
		size    = 100   // [-size..+size]の範囲の画像キャンパス
		nframes = 64    // アニメーションフレーム数
		delay   = 8     // フレーム間遅延、単位は10ms
	)
	freq := rand.Float64() * 3.0 // yオシレータの相対周波数
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // 位相差
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			//img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
			// 線を緑にする
			img.Set(size+int(x*size+0.5), size+int(y*size+0.5), color.RGBA{0x00, 0xff, 0x00, 0xff})
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	//標準出力だと文字化けして正しく表示されなかった
	//gif.EncodeAll(out, &anim) // 注意: エンコードエラーを無視
	//ファイル出力
	f, _ := os.OpenFile("data/lissajous_ex5.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, &anim)
}
