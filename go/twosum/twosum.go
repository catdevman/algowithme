package twosum

func TwoSum(nums []int, target int) []int {
	// add a comment to trigger workflow
	out := []int{}
outerLoop:
	for i := 0; i <= len(nums)-2; i++ {
		for j := i + 1; j <= len(nums)-1; j++ {
			if nums[i]+nums[j] == target {
				out = append(out, i)
				out = append(out, j)
				break outerLoop
			}
		}
	}
	return out
}
