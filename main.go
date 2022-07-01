package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello, World!")
	rand.Seed(time.Now().UnixMicro())
	buildUniverse(100, 20, 10)
}
