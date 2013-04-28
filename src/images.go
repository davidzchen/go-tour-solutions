package main

import (
    "code.google.com/p/go-tour/pic"
    "image"
    "image/color"
)

func F0(x, y int) uint8 {
    return uint8(x * y)
}

func F1(x, y int) uint8 {
    return uint8((x + y) / 2)
}

func F2(x, y int) uint8 {
    for i := 0; i < y; i++ {
        x *= x
    }
    return uint8(x)
}

type Image struct{
    Width int
    Height int
    Function func (x, y int) uint8
}

func (img Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, img.Width, img.Height)
}

func (img Image) At(x, y int) color.Color {
    return color.RGBA{img.Function(x, y), img.Function(x, y), 255, 255}
}

func main() {
    m := Image{500, 500, F0}
    pic.ShowImage(m)
}
