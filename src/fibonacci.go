package main

import "fmt"

// fibonacci is a function that returns a function that returns an int
func fibonacci() func() int {
    f0 := -1
    f1 := 0
    return func() int {
        if f0 == -1 {
            f0++
            return f0
        } else if f1 == 0 {
            f1++
            return f1
        }

        last := f1
        f1 = f0 + f1
        f0 = last
        return f1
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
