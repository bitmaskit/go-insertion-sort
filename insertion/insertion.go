package insertion

func Sort(nums []int) {
	for j := 1; j < len(nums); j++ {
		ins := nums[j]
		i := j - 1
		for ; i >= 0 && ins < nums[i]; i-- {
			nums[i+1] = nums[i]
		}
		nums[i+1] = ins
	}
}