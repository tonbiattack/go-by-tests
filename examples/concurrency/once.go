package concurrencyexample

import "sync"

func CallsAfterOncePanic() (calls int) {
	var once sync.Once
	run := func() {
		defer func() {
			_ = recover()
		}()
		once.Do(func() {
			calls++
			panic("first call failed")
		})
	}
	run()
	run()
	return calls
}
