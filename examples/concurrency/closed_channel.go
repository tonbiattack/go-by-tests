package concurrencyexample

func ReadAfterClose() (first int, afterClose int, open bool) {
	values := make(chan int, 1)
	values <- 7
	close(values)

	first = <-values
	afterClose, open = <-values
	return first, afterClose, open
}
