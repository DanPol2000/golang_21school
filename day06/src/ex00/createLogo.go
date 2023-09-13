package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	const width, height = 300, 300

	logo := image.NewNRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			logo.Set(x, y, color.NRGBA{
				R: 0,
				G: 0,
				B: 0,
				A: 0,
			})
			radius := (width/2-x)*(width/2-x) + (height/2-y)*(height/2-y)
			if radius < 9000 {
				logo.Set(x, y, color.NRGBA{
					R: 255,
					G: uint8(255 - 255*radius/9000),
					B: 0,
					A: 255,
				})
			}
		}
	}

	f, err := os.Create("logo.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, logo); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
