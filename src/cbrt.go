package main

import (
    "fmt"
    "math/cmplx"
)

const Delta float64 = 0.0000000001

func Cbrt(x complex128) complex128 {
    z := x / 2
    z0 := complex128(0)
    for {
        z = z - (cmplx.Pow(z, 3) - x) / (3 * z * z)
        if cmplx.Abs(z - z0) < Delta {
            break
        }
        z0 = z
    }

    return z
}

func main() {
    fmt.Println(Cbrt(2))
}
