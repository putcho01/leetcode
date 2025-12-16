
// https://leetcode.com/problems/shuffle-the-array/?envType=problem-list-v2&envId=dsa-linear-shoal-array-i
func shuffle(nums []int, n int) []int {
	ans := make([]int, 0, len(nums))
	for i := 0; i < n; i++ {
		ans = append(ans, nums[i])
    ans = append(ans, nums[n+i])
	}
	return ans
}