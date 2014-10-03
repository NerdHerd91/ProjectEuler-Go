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

    sum := sumOfSquares(n)
    square := squareOfSums(n)
    fmt.Printf("Sum of the squares for first %d natural numbers: %d\n", n, sum)
    fmt.Printf("Square of the sum of the first %d natural numbers: %d\n", n, square)
    fmt.Printf("Sum-Square difference: %d\n", square - sum)
}

func sumOfSquares(n int) int {
    sum := 0
    for i := 1; i <= n; i++ {
        sum += i * i
    }
    return sum
}

func squareOfSums(n int) int {
    sum := 0
    for i := 1; i <= n; i++ {
        sum += i
    }
    return sum * sum
}
