package arrays

func Sum(numbers []int) int {
	var sum int
	for _, num := range numbers {
		sum += num
	}
	return sum;
}
func SumAllTails(toSum ...[]int) []int{
	var sums []int
	for _, nums := range toSum {
		sums = append(sums, Sum(nums[1:]))
	}
	return sums
}
