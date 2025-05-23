package main

type student struct {
	name  string
	score int
}

func scoreAvg(scores []float64) (result float64) {
	for _, score := range scores {
		result += score
	}
	if len(scores) == 0 {
		return 0
	}
	return result / float64(len(scores))
}

func main() {

}
