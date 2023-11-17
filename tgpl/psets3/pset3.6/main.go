// Pset3.6 supersampling.
package main

import (
	"image"
	"image/color"
	"image/png"
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
		for px := 0; px < width; px++ {
			// Image point (px, py) represents complex value z.
			var red, green, blue, alpha uint32
			var cnt uint32
			for i := 0; i < 2; i++ {
				for j := 0; j < 2; j++ {
					a := px + i
					b := py + j
					if a <= width && b <= height {
						cnt++
						ax := float64(a)/width*(xmax-xmin) + xmin
						by := float64(b)/height*(xmax-xmin) + xmin
						s := complex(ax, by)
						sr, sg, sb, sa := mandelbrot(s).RGBA()
						red += sr
						green += sg
						blue += sb
						alpha += sa
					}
				}
			}
			img.Set(px, py, color.RGBA{uint8(red / cnt), uint8(green / cnt), uint8(blue / cnt), uint8(alpha / cnt)})
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
