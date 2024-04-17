package section1

// swap the value of a and b by assigning
func SwapVariableByReassigning(a, b int) (int, int) {
	b, a = a, b
	return a, b
}

func SwapVariableWithReturn(a, b int) (int, int) {
	return b, a
}
