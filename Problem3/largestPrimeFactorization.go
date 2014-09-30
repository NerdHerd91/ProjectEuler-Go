package main

import (
    "fmt"
    "log"
    "os"
    "strconv"
)

func main() {
    if len(os.Args) < 2 {
        log.Fatal("An integer N must be provided as the first argument.")
    }

    // Attempt to parse the argument for n.
    n, err := strconv.Atoi(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(findLargestPrimeFactor(n))
}

func findLargestPrimeFactor(value int) (int) {
    nextPrime := 2
    for i := 2; i < value / i; i++  {
        for value % i == 0 {
            if nextPrime != i {
                nextPrime = i
            }
            value /= i
        }
    }
    if value > 1 {
        nextPrime = value
    }
    return nextPrime
}
