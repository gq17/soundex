package soundex

import (
	"math/rand"
)

type entry struct{}

func Random(n int) int {
	return rand.Intn(n)
}
