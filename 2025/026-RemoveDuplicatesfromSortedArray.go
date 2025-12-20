// https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150

// 制約: -100 <= nums[i] <= 100

func removeDuplicates(nums []int) int {
    previousNum := -101
    k :=0
    for _, val := range nums {
        if val != previousNum {
            nums[k] = val
            previousNum = val
            k++
        }
    }
    return k
}