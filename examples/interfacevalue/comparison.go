package interfacevalue

func CompareSlicesAsAnyPanics() (recovered any) {
	left := any([]int{1})
	right := any([]int{1})
	defer func() {
		recovered = recover()
	}()
	_ = left == right
	return nil
}
