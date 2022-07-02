package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello, World!")
	rand.Seed(time.Now().UnixMicro())

	worldX = 50
	worldY = 50

	s := worldX * worldY

	p := int(float64(s) * 0.33)
	st := int(float64(s) * 0.05)
	buildUniverse(s, p, st, int(md))

	d := distance(0, 0, s, s)
	fmt.Println(d)
}
