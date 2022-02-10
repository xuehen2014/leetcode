package BinarySearch

/*
 * https://leetcode-cn.com/problems/binary-search/
 */

//二分查找--迭代
func search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	begin := 0
	end := length
	for begin < end { //左闭右开区间 [begin, end), begin==end区间为空区间
		mid := begin + (end-begin) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			begin = mid +1
		} else {
			end = mid
		}
	}
	return -1
}

//二分查找--递归
func search2(nums []int, target int) int {
	left, right := 0, len(nums)
	return binarySearch(nums, target, left, right)
}

func binarySearch(nums []int, target int, left, right int) int {
	if left >= right { //左闭右开区间 [begin, end), begin==end区间为空区间
		return -1
	}
	middle := left + (right-left)/2
	middleNum := nums[middle]
	if middleNum == target {
		return middle
	} else if middleNum < target {
		return binarySearch(nums, target, middle+1, right)
	} else {
		return binarySearch(nums, target, left, middle)
	}
}





