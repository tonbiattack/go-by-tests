package concurrencyexample

func SendAfterClosePanics() (recovered any) {
	values := make(chan int)
	close(values)

	defer func() {
		recovered = recover()
	}()
	values <- 1
	return nil
}
