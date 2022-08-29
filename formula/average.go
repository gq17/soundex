package formula

import (
	"math"
	"math/rand"
	"time"
)

// n个相互独立的均匀分布的随机数的和

var ra1, ra2, ra3, ra4 rand.Rand

func initra1234() {
	ra1 = *rand.New(rand.NewSource(time.Now().Unix()))
	ra2 = *rand.New(rand.NewSource(time.Now().Unix()))
	ra3 = *rand.New(rand.NewSource(time.Now().Unix()))
	ra4 = *rand.New(rand.NewSource(time.Now().Unix()))
}

func averageRandom() float64 {
	var a float64

	a += ra1.Float64()
	a += ra2.Float64()
	a += ra3.Float64()
	a += ra4.Float64()

	a -= 2
	a *= math.Sqrt(3)

	return a
}
