package watermark

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

func main() {
	imgb, _ := os.Open("image.jpg")
	img, _ := jpeg.Decode(imgb)
	defer imgb.Close()

	wmb, _ := os.Open("watermark_opacity.png")
	watermark, _ := png.Decode(wmb)
	defer wmb.Close()

	offset := image.Pt(0, 0)
	b := img.Bounds()
	m := image.NewRGBA(b)
	draw.Draw(m, b, img, image.ZP, draw.Src)
	draw.Draw(m, watermark.Bounds().Add(offset), watermark, image.ZP, draw.Over)

	imgw, _ := os.Create("watermarked_0_0.jpg")
	jpeg.Encode(imgw, m, &jpeg.Options{jpeg.DefaultQuality})
	defer imgw.Close()
}
