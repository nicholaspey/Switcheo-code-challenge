package main

import "fmt"

// Complexity = O(n). Efficient and linear time complexity, constant space complexity
func sum_to_n_a(n int) int {
    sum := 0
    for i := 1; i <= n; i++ {
        sum += i
    }
    return sum
}

// Complexity = O(1). Most efficient, constant time and space complexity
func sum_to_n_b(n int) int {
    return n * (n + 1) / 2
}

// Complexity = O(n). Less efficient, due to recusion overhead, linear time complexity, linear space complexity. Hence worst
func sum_to_n_c(n int) int {
    if n == 0 {
        return 0
    }
    return n + sum_to_n_c(n-1)
}

func main() {
    // Test values
    n := 5
    
    // Call the functions and print the results
    fmt.Printf("sum_to_n_a(%d) = %d\n", n, sum_to_n_a(n))
    fmt.Printf("sum_to_n_b(%d) = %d\n", n, sum_to_n_b(n))
    fmt.Printf("sum_to_n_c(%d) = %d\n", n, sum_to_n_c(n))
}