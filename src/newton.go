package main

import (
    "fmt"
    "math"
)

const Delta float64 = 0.0000000001

func Sqrt(x float64) float64 {
    if x == 0 {
        return 0
    }

    z := x / 2
    z0 := 0.0
    for {
        z = z - (z * z - x) / (2 * z)
        if math.Abs(z - z0) < Delta {
            break
        }
        z0 = z
    }

    return z
}

func main() {
    for i := 0; i < 40; i++ {
        estimate := Sqrt(float64(i))
        actual := math.Sqrt(float64(i))
        fmt.Print(i)
        fmt.Print("\t")
        fmt.Print(estimate)
        fmt.Print("\t")
        fmt.Print(actual)
        fmt.Print("\t")
        fmt.Println(math.Abs(estimate - actual))
    }
}
