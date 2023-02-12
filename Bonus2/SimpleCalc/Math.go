package SimpleCalc

type Pair struct {
	first  int
	second int
}

func (P *Pair) add() int {
	return P.first + P.second
}

func (P *Pair) subtract() int {
	return P.first - P.second
}

func (P *Pair) multiply() int {
	return P.first * P.second
}
