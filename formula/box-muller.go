package formula

import (
	"math"
	"math/rand"
	"time"
)

const PI float64 = 3.1415926

var rb1, rb2 rand.Rand

func initrb12() {
	rb1 = *rand.New(rand.NewSource(time.Now().Unix()))
	rb2 = *rand.New(rand.NewSource(time.Now().Unix()))
}

func box_muller() float64 {

	x, y := rb1.Float64(), rb2.Float64()

	a := math.Sqrt(-2 * math.Log(x))
	b := 2 * PI * y
	c := a * math.Cos(b)

	return c
}
