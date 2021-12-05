package core

type FibonacciValue struct {
	Number int
	Value  int
}

func Fibonacci(n int) []FibonacciValue {
	var res []FibonacciValue
	a, b := 0, 1
	for i := 0; i < n; i++ {
		temp := a
		a = b
		b = temp + a
		res = append(res, FibonacciValue{
			Number: i + 1,
			Value:  a,
		})
	}

	return res
}
