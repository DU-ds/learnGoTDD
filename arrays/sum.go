package arrays

func Sum(numbers []int) int {
	var sum int
	for _, num := range numbers {
		sum += num
	}
	return sum;
}
func SumAll(toSum ...[]int) []int{
	var sums []int
	for _, nums := range toSum {
		sums = append(sums, Sum(nums))
	}
	return sums
}
