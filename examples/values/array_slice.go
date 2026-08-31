package values

func ChangeArray(values [2]string) [2]string {
	values[0] = "changed"
	return values
}

func ChangeSlice(values []string) {
	values[0] = "changed"
}
