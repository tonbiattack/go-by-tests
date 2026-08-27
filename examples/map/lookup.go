package mapexample

func Lookup(scores map[string]int, name string) (int, bool) {
	score, ok := scores[name]
	return score, ok
}
