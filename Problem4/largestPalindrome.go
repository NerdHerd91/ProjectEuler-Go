package main

import (
    "fmt"
    "math"
    "strconv"
)

func main() {
    fmt.Println(findLargestPalindrome())
}

func findLargestPalindrome() (int)  {
    maxPal := 0
    for i := 100; i <= 999; i++ {
        for j := 100; j <= 999; j++ {
            value := i * j
            if isPalindrome(value) {
                maxPal = int(math.Max(float64(maxPal), float64(value)))
            }
        }
    }
    return maxPal
}

func isPalindrome(value int) (bool) {
    strVal := strconv.Itoa(value)
    reverse := ""
    for _, r := range strVal {
        reverse = string(r) + reverse
    }
    return strVal == reverse
}
