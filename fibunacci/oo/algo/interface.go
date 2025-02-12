package algo

type FibGen interface {
	Generate(nrOfElements int) (result []int)
}
