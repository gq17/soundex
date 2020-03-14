package soundex

var (
	dist = NewDistribution()
)

func GetIntNumber() int {
	return dist.GetIntNumber()
}

func GetFloatNumber() float32 {
	return dist.GetFloatNumber()
}
