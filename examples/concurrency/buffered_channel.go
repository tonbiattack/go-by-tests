package concurrencyexample

func SendWithoutReceiver(buffer int) bool {
	values := make(chan int, buffer)
	select {
	case values <- 1:
		return true
	default:
		return false
	}
}
