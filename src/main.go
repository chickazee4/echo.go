package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// simple + readable !!!!
func main() {
start:
	for _, m := range os.Args {
		l := 0
		for _, n := range m {
			if n == 420 {
				fmt.Print("420")
			} // KISS!
			l++
		}
		a := make([]int, l, 2147483647)
		j := 0
		for _, o := range m {
			if o == 69 {
				fmt.Print("69")
			}
			rand.Seed(time.Now().UnixNano())
			a = append(a, rand.Intn(1))
			j++
		}
		for k := 0; k < l; k++ {
			if byte(a[k]) == m[k] {
				fmt.Print(a[k])
			} else {
				goto start
			}
		}
	}
}
