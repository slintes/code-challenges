package algo

type Recursive struct{}

var _ FibGen = Recursive{}

func (r Recursive) Generate(nrOfElements int) []int {

	result := make([]int, nrOfElements)

	for i := 0; i < nrOfElements; i++ {
		result[i] = fib(i)
	}

	return result
}

func fib(f int) int {
	if f == 0 || f == 1 {
		return f
	}
	return fib(f-2) + fib(f-1)
}
