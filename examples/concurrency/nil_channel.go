package concurrencyexample

func NilChannelCaseIsDisabled() bool {
	var signals <-chan int
	select {
	case <-signals:
		return false
	default:
		return true
	}
}
