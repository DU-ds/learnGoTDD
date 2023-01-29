package arrays

func Sum(numbers []int) int {
	var sum int
	for _, num := range numbers {
		sum += num
	}
	return sum;
}
func SumAll(toSum ...[]int) []int{
	var n int = len(toSum)
	var sums []int //= make([]int, n)
	for i := 0; i < n; i++ {
		sums = append(sums, Sum(toSum[i]))
	}
	return sums
}
