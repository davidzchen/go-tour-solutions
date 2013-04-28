package main

import "code.google.com/p/go-tour/pic"

const (
    Height = 500
    Width = 500
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

func Pic(dx, dy int) [][]uint8 {
    picture := make([][]uint8, Width)
    for i := 0; i < Width; i++ {
        picture[i] = make([]uint8, Height)
        for j := 0; j < Height; j++ {
            picture[i][j] = F0(i, j)
        }
    }

    return picture
}

func main() {
    pic.Show(Pic)
}
