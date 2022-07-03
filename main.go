package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello, World!")
	rand.Seed(time.Now().UnixMicro())

	worldX = 50
	worldY = 50

	s := worldX * worldY
	md := float64(math.Sqrt(float64(s))) / 6.0

	p := int(float64(s) * 0.33)
	st := int(float64(s) * 0.05)
	buildUniverse(s, p, st, int(md))

}
