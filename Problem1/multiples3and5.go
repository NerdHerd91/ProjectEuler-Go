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

    fmt.Println(computeSum(n))
}

func computeSum(limit int) (int) {
    sum := 0
    for i := 0; i < limit; i++ {
        if i % 3 == 0 || i % 5 == 0 {
            sum += i
        }
    }
    return sum
}
