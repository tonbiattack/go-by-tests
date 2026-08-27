package stringexample

import "strings"

func CopyAfterWritePanics() (recovered any) {
	var builder strings.Builder
	builder.WriteString("go")
	copied := builder

	defer func() {
		recovered = recover()
	}()
	copied.WriteString("tests")
	return nil
}
