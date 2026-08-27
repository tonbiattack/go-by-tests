package loopexample

func VisitWhileAppending() (visited []int, final []int) {
	values := []int{1, 2, 3}
	for _, value := range values {
		visited = append(visited, value)
		values = append(values, value+10)
	}

	return visited, values
}
