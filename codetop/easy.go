package codetop

//////////////////////////////////////////////////////////////////////////////////////////////
// leetcode No.1
// 思路是：用一个 map 维护数组已存在的数和坐标，遍历数组，看map中是否有 与当前数组之和是target的。
/////////////////////////////////////////////////////////////////////////////////////////////
func twoSum(nums []int, target int) []int {
	var exist = make(map[int]int)
	for i, num := range nums {
		if v, ok := exist[target-num]; ok {
			return []int{i, v}
		}
		exist[num] = i
	}
	return []int{}
}
