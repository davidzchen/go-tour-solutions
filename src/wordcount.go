package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    words := strings.Fields(s)
    counts := make(map[string]int)
    for i := 0; i < len(words); i++ {
        if _, ok := counts[words[i]]; ok == false {
            counts[words[i]] = 1
        } else {
            counts[words[i]]++
        }
    }

    return counts
}

func main() {
    wc.Test(WordCount)
}
