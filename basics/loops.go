package main

func loops() int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}