package hash

import "sort"

func threeSum(nums []int) [][]int {
	//double points
	if nums == nil || len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	res := make([][]int, 0)
	//第一个循环
	for i := 0; i < len(nums); i++ {
		num1 := nums[i]
		if num1 > 1 && num1 == nums[i-1] {
			continue
		}
		if num1 > 0 {
			break
		}
		//第二个循环,移动左右指针
		left, right := i+1, len(nums)-1
		for left < right {
			num2, num3 := nums[left], nums[right]
			sum := num1 + num2 + num3
			if sum == 0 {
				res = append(res, []int{num1, num2, num3})
				//去重逻辑
				for left < right && nums[left] == num2 {
					left++
				}
				for left < right && nums[right] == num3 {
					right--
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}

func fourSum(nums []int, target int) [][]int {
	if nums == nil || len(nums) < 4 {
		return nil
	}
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-3; i++ {
		num1 := nums[i]
		//过滤重复数据
		if i > 0 && num1 == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			num2 := nums[j]
			if j > i+1 && num2 == nums[j-1] {
				continue
			}
			left, right := j+1, len(nums)-1
			for left < right {
				num3, num4 := nums[left], nums[right]
				sum := num1 + num2 + num3 + num4
				if sum == target {
					res = append(res, []int{num1, num2, num3, num4})
					for left < right && nums[left] == num3 {
						left++
					}
					for left < right && nums[right] == num4 {
						right--
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}

	}
	return res
}
func smallestRangeII(nums []int, k int) int {
	sort.Ints(nums)
	if nums[len(nums)-1] < k {

	}
	min, max := nums[0], nums[len(nums)-1]
	res := max - min
	for i := 0; i < len(nums)-1; i++ {
		a, b := nums[i], nums[i+1]
		res = min(res, max(ma-k, a+k)-min(mi+k, b-k))
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
