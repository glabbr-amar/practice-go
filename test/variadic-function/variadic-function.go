package main

func add(num ...int) int {
	sum := 0
	for _, data := range num {
		sum += data
	}
	return sum
}
