package main

import (
	"fmt"

	"github.com/slintes/code-challenges/fibunacci/oo/algo"
)

func main() {

	nrOfElements := 10

	algos := []algo.FibGen{algo.Loop{}, algo.Recursive{}}

	for _, algo := range algos {
		fmt.Printf("%v\n", algo.Generate(nrOfElements))
	}

}
