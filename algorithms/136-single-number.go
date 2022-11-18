package main

//https://leetcode.com/problems/single-number/
//https://leetcode.com/problems/single-number/discuss/559971/PythonJSJavaGoC%2B%2B-O(n)-by-XOR-w-Hint
func singleNumber(nums []int) int {
	m := map[int]bool{}
	for _, num := range nums {
		if _, found := m[num]; found == true {
			delete(m, num)
		} else {
			m[num] = true
		}
	}
	for k := range m {
		return k
	}
	return -1
}
