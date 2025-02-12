package algo

type Loop struct{}

var _ FibGen = Loop{}

func (l Loop) Generate(nrOfElements int) []int {

	result := make([]int, nrOfElements)

	f1, f2 := 0, 1
	for i := 0; i < nrOfElements; i++ {
		result[i] = f1
		sum := f1 + f2
		f1 = f2
		f2 = sum
	}

	return result
}
