package mapexample

func ReadNilMap() int {
	var scores map[string]int
	return scores["missing"]
}

func WriteNilMapPanics() (recovered any) {
	var scores map[string]int
	defer func() {
		recovered = recover()
	}()
	scores["go"] = 1
	return nil
}
