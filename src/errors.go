package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %f\n", float64(e))
}

const Delta float64 = 0.0000000001

func Sqrt(x float64) (float64, error) {
    if x < 0 {
        return 0, ErrNegativeSqrt(x)
    }

    if x == 0 {
        return 0, nil
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

    return z, nil
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}
