package twosum

func TwoSum(nums []int, target int) []int {
	out := []int{}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]+nums[i+1] == target {
			out = append(out, i)
			out = append(out, i+1)
			break
		}
	}
	return out
}
