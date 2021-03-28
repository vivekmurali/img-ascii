package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/disintegration/imaging"
	"github.com/nfnt/resize"
)

func main() {
	src, err := imaging.Open("logo_250.png")
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	img3 := imaging.Transpose(src)

	img := resize.Resize(50, 0, img3, resize.Lanczos3)

	I := image.NewGray(img.Bounds())
	sum := 0
	for i := 0; i < img.Bounds().Max.X; i++ {
		for j := 0; j < img.Bounds().Max.Y; j++ {
			o := img.At(i, j)
			n := color.GrayModel.Convert(o)
			I.Set(i, j, n)
			sum += int(I.GrayAt(i, j).Y)
		}
	}
	mean := sum / (I.Bounds().Max.X * I.Bounds().Max.Y)

	for i := 0; i < I.Bounds().Max.X; i++ {
		for j := 0; j < I.Bounds().Max.Y; j++ {
			if int(I.GrayAt(i, j).Y) >= mean {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}

		}
		fmt.Printf("\n")
	}

}
