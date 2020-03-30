package soundex

type Distribution struct {
	dtype string
	seed  string
}

func NewDistribution(t string) *Distribution {
	return &Distribution{
		dtype: t,
	}
}

func (dist *Distribution) GetGaussianIntNumber() int {
	return 0
}

func (dist *Distribution) GetFloatNumber() float32 {
	return 0
}
