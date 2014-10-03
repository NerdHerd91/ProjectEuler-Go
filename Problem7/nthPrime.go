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

    fmt.Println(computeNthPrime(n))
}

func computeNthPrime(n int) int {
    primeCounter := 1
    prime := 2
    for  i := 3 ; primeCounter < n; i += 2 {
        if isPrime(i) {
            primeCounter++
            prime = i
        }
    }
    return prime
}

func isPrime(n int) bool {
    for i:= 2; i < n; i++ {
        if n % i == 0 {
            return false
        }
    }
    return true
}
