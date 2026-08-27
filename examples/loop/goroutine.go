package loopexample

import "sort"

func RangeValuesFromGoroutines() []int {
	values := make(chan int, 3)
	for index := range 3 {
		go func() {
			values <- index
		}()
	}

	got := []int{<-values, <-values, <-values}
	sort.Ints(got)
	return got
}
