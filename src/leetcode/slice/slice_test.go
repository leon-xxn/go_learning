package slice

import "testing"

func firstMissingPositive(nums []int) int {
	// hash
	max := len(nums)
	for i, v := range nums {
		if v == 0 {
			nums[i] = -1
		}
	}
	for _, v := range nums {
		if v > 0 && v <= max {
			nums[v-1] = 0
		}
	}
	num := 0
	for k, v := range nums {
		if v == 0 {
			if num > k {
				num = k
			}
			continue
		} else if v == max+1 {
			continue
		}
		return k + 1
	}
	if num == 0 {
		num = max + 1
	}
	return num
}
func TestFirstMissingPositive(t *testing.T) {
	res := firstMissingPositive([]int{3, 4, -1, 1})
	t.Log(res)
}
