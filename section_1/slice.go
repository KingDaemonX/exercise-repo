package section1

func SumEvenNumberInSlice(nums []int) int {
	result := 0

	for _, num := range nums {
		if num%2 == 0 {
			result += num
		}
	}

	return result
}
