package main

import (
	"fmt"
	"math"
)

func main() {
	a := []float64{
		1.2,
		2.3,
	}
	fmt.Println(Gps(10, a))
}
func Gps(s int, x []float64) int {
	var max float64 = 0
	for i := 0; i < len(x)-1; i++ {
		speed := (float64(3600) * (x[i+1] - x[i])) / float64(s)
		max = math.Max(max, speed)
	}
	return int(max)
}
