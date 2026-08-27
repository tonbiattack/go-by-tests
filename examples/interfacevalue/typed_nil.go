package interfacevalue

import "io"

func IsNilReader(reader io.Reader) bool {
	return reader == nil
}
