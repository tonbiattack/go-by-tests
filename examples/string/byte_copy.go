package stringexample

func ChangeFirstByte(original string) (unchanged string, changed string) {
	bytes := []byte(original)
	bytes[0] = 'G'
	return original, string(bytes)
}
