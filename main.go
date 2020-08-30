package main

import (
	"fmt"
	"log"

	"github.com/harukitosa/box/mat"
)

func main() {
	i := []float64{1, 2, 3, 4, 10, 1}
	fmt.Println(i)
	m, err := mat.NewMat(i, 2, 3)
	if err != nil {
		log.Fatal(err)
	}
	m.Print()
	j := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	n, err := mat.NewMat(j, 3, 3)
	if err != nil {
		log.Fatal(err)
	}
	n.Print()
}
