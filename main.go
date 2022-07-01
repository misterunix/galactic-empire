package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello, World!")
	rand.Seed(time.Now().UnixMicro())
	s := 1000
	p := int(float64(s) * 0.33)
	st := int(float64(s) * 0.05)
	buildUniverse(s, p, st)
}
