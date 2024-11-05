package calc

type Addition struct{}
type Subtract struct{}
type Multiply struct{}
type Division struct{}

func (this Addition) Calculate(a, b int) int {
	return a + b
}

func (this Subtract) Calculate(a, b int) int {
	return a - b
}

func (this Multiply) Calculate(a, b int) int {
	return a * b
}

func (this Division) Calculate(a, b int) int {
	return a / b
}
