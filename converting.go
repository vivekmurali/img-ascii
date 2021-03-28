package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	f, err := os.Open("logo_250.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	I := image.NewGray(img.Bounds())
	for i := 0; i < img.Bounds().Max.X; i++ {
		for j := 0; j < img.Bounds().Max.Y; j++ {
			o := img.At(i, j)
			n := color.GrayModel.Convert(o)
			I.Set(i, j, n)
		}
	}
	fmt.Println(I)
	nf, err := os.Create("test_logo.png")
	defer nf.Close()
	err = png.Encode(nf, I)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(img)
}
