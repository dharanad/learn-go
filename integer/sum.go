package integer

func Sum(nums []int) int {
	var sum int
	for _, val := range nums {
		sum += val
	}
	return sum
}

func SumAll(numsArr ...[]int) []int {
	var sums []int
	for _, arr := range numsArr {
		sums = append(sums, Sum(arr))
	}
	return sums
}

func SumAllTails(numsArr ...[]int) []int {
	var sums []int
	for _, arr := range numsArr {
		if len(arr) > 1 {
			sums = append(sums, Sum(arr[1:]))
		} else {
			sums = append(sums, 0)
		}
	}
	return sums
}
