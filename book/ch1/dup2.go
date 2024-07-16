// Dup2 prints the count and text lines that appear more than once
// in the input.  It reads from stdin or from a list of names files.
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    counts := make(map[strings]int)
    files := os.Args[1:]
    if len(files) == 0 {
        countLines(os.Stdin, counts)
    } else {
        for _, arg :=range files {
            f, err := os.Open(arg)
	    if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
		continue
	    }
	    countLines(f, counts)
	    f.Close()
	}
    }
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
	}
    }
}
    func countLines(f *os.File, counts map[string]int) {
        input := bufio.NewScanner(f)
	for input.Scan() {
	    count[input.Text()]++
    }
    // NOTE: ignoring potential errors form input.Err()
}
