package soundex

type Distribution struct{
	dtype string
}

func NewDistribution() *Distribution {
	return &Distribution{}
}

func (dist *Distribution) GetIntNumber() int {
	return 0
}

func (dist *Distribution) GetFloatNumber() float32 {
	return 0
}
