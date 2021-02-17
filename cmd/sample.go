package main

import (
	"fmt"
	"math/rand"
)

// type Source interface {
//     Int63() int64
//     Seed(seed int64)
// }

func main() {
	// rand_int = rand.Intn(10)
	a := rand.New(rand.NewSource(1))
	for i:=0; i < 10; i ++ {
		fmt.Println(i)
	}
	fmt.Println(*a)
}
