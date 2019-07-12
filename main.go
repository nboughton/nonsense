package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/nboughton/swnt/content/name"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	l, m, n := 0, 0, 0
	flag.IntVar(&n, "n", 10, "Number of nonsense words to generate")
	flag.IntVar(&l, "l", 10, "Maxlength of generated words")
	flag.IntVar(&m, "m", 4, "Minlength of generated words")
	flag.Parse()

	for i := 0; i < n; i++ {
		ln := rand.Intn(l)
		if ln < m {
			ln = m
		}

		fmt.Println(name.Generate(ln))
	}
}
