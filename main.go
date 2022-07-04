package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// globals
var sectorCount int

func main() {
	fmt.Println("Hello, World!")
	rand.Seed(time.Now().UnixMicro())

	worldX = 10
	worldY = 10

	sectorCount = worldX * worldY
	md := float64(math.Sqrt(float64(sectorCount))) / 6.0

	p := int(float64(sectorCount) * 0.33)
	st := int(float64(sectorCount) * 0.05)
	buildUniverse(p, st, int(md))

}
