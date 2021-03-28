package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"math/rand"
	"os"
)

func main() {
	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}
	white
	img := image.NewGray(image.Rect(0, 0, 100, 100))
	for i := 0; i < 100; i++ {
		x := rand.Intn(101)
		y := rand.Intn(101)
		img.Set(x, y, color.Gray{200})
	}
	fmt.Printf("img = %T\n", img)
	//fmt.Println(img)
	err = png.Encode(f, img)
	if err != nil {
		log.Fatal(err)
	}
	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}

}
