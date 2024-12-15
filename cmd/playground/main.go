package main

// #include <stdio.h>
//
// void hello() {
//    printf("Hello, C\n");
//}
//
import "C"

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"

	"golang.design/x/clipboard"
)

func main() {
	fmt.Println("Hello, Go")
	C.hello()

	if err := mainE(); err != nil {
		panic(err.Error())
	}
}

func mainE() error {
	var buf bytes.Buffer
	if err := draw(&buf, 1024, 768); err != nil {
		return err
	}

	_ = clipboard.Write(clipboard.FmtImage, buf.Bytes())

	return nil
}

func draw(w io.Writer, width int, height int) error {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for h := 0; h < height; h++ {
		b := 255 - uint8(255.0*(float64(h)/float64(height)))

		for w := 0; w < width; w++ {
			img.Set(w, h, color.RGBA{0, 0, b, 0xFF})
		}
	}

	return png.Encode(w, img)
}
