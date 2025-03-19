package ways_to_split

func WaysToSplit(nums []int) int {
	count := 0
	size := len(nums)
	for i := 1; i < size; i++ {
		nums[i] += nums[i-1]
	}
	for i := 1; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			left := nums[i-1]
			mid := nums[j-1] - nums[i-1]
			right := nums[size-1] - nums[j-1]
			if left <= mid && mid <= right {
				count += 1
			}
		}
	}
	return count
}
