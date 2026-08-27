package deferflow

func CapturedArgument() (events []string) {
	events = []string{}
	label := "before"
	defer func(value string) {
		events = append(events, value)
	}(label)
	label = "after"
	return
}

func LIFO() (events []int) {
	defer func() { events = append(events, 1) }()
	defer func() { events = append(events, 2) }()
	return events
}
