package soundex

import (
	"math"
	"math/rand"
	"time"
)

const (
	GAUSSIAN = "gaussian"
)

var (
	dist = NewDistribution(GAUSSIAN)
)

func GetIntGaussianNumber() int {
	return dist.GetIntGaussianNumber()
}

func GetExponentialNumber(lambda float64) float64 {
	if lambda < 0 {
		return -1
	}

	return (-1.0 / lambda) * math.Log(rand.Float64())
}

func GetFloatGaussianNumber(mean float64, cov float64) float64 {
	m, n := rand.Float64(), rand.Float64()
	r := math.Sqrt(-2 * math.Log(n))
	t := 2 * math.Pi * m

	return r * math.Cos(t)
}

func ResetSeed() {
	rand.Seed(time.Now().Unix())
}
