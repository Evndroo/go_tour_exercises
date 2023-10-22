package main

// Fibonacci is a function that returns
// a function that returns an int.
func Fibonacci() func() int {
	previous, current := 0, 0

	return func() int {
		if current == 0 && previous == 0 {
			current = 1
			return previous
		} else {
			previous, current = current, previous+current
			return previous
		}
	}
}
