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

    fmt.Println(computeFibonacciSum(1, 2, n))
}

func computeFibonacciSum(f1, f2, limit int) (int) {
    fibonacciSum := 0
    if (f1 % 2 == 0 && f2 <= limit) {
        fibonacciSum += f1
    }
    if (f2 % 2 == 0 && f2 <= limit) {
        fibonacciSum += f2
    }

    for ; true;  {
        sum := f1 + f2
        f1 = f2
        f2 = sum

        if (f2 > limit) {
            break
        } else if (f2 % 2 == 0) {
            fibonacciSum += f2
        }
    }
    return fibonacciSum
}
