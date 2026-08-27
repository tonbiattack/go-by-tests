package deferflow

func RecoverResults() (directlyRecovered bool, helperRecovered bool, panicEscaped bool) {
	directlyRecovered = recoverDirectly()

	func() {
		defer func() {
			panicEscaped = recover() != nil
		}()
		helperRecovered = recoverThroughHelper()
	}()

	return directlyRecovered, helperRecovered, panicEscaped
}

func recoverDirectly() (recovered bool) {
	defer func() {
		recovered = recover() != nil
	}()
	panic("direct recovery")
}

func recoverThroughHelper() (recovered bool) {
	defer func() {
		recovered = recoverFromHelper() != nil
	}()
	panic("indirect recovery")
}

func recoverFromHelper() any {
	return recover()
}
