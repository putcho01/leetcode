https://leetcode.com/problems/remove-element/?envType=study-plan-v2&envId=top-interview-150

func removeElement(nums []int, val int) int {
    k := 0
    for _, v := range nums {
        if v != val {
            nums[k] = v
            k++
        }
    }
    return k
}