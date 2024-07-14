// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var s, sep string
	start := time.Now()
	for i := 1; i < len(os.Args); i++ {
		s += sep + "[" + fmt.Sprint(i) + "]" + os.Args[i]
		sep = " "
	}
	fmt.Println(os.Args[0])
	fmt.Println(s)
	elapsed := time.Since(start).Milliseconds()
	fmt.Printf("%dms elapsed\n", elapsed)
}
