package main

import (
    "os"
    "testing"
)

// BenchmarkMain benchmarks the main function
func BenchmarkMain(b *testing.B) {
    // Set up arguments for the benchmark
    os.Args = []string{"echo1.go", "Hello", "World"}

    // Reset the timer to exclude setup time from the benchmark
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        main()
    }
}

