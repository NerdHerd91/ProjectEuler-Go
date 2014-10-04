package main

import (
    "fmt"
    "log"
    "math"
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

    fmt.Println(sumPrimes(n))
}

func sumPrimes(n int) int {
    sum := 2
    for i := 3; i < n; i += 2 {
        if isPrime(i) {
            sum += i
        }
    }
    if n < 3 {
        return 0
    }
    return sum
}

func isPrime(n int) bool {
    for i := 3; i <= int(math.Ceil(math.Sqrt(float64(n)))); i += 2 {
        if n % i == 0 {
            return false
        }
    }
    return true
}
