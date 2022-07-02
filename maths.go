package main

import "math"

// return the distance between two sectors
func distance(s1, s2 int) int {
	x1, y1 := sectorToCoord(s1)
	x2, y2 := sectorToCoord(s2)
	return int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
}

// convert a sector into a x,y coordinate
// returns the x,y coordinate of the sector
func sectorToCoord(sector int) (int, int) {
	x := sector % worldX
	y := sector / worldX
	return x, y
}
