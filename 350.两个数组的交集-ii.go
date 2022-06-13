/*
 * @Author: Y.t
 * @Date: 2022-06-11 17:03:21
 * @LastEditors: Y.t
 * @LastEditTime: 2022-06-11 22:56:36
 * @Description:
 */
/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {

	//返回数组
	var res []int
	//方案一 位置排序
	// //创建字典来存储数组1
	// mCache := map[int]int{}
	// for _, v := range nums1 {
	// 	mCache[v] += 1
	// }

	// //遍历数组2判定重复数字
	// for _, v := range nums2 {
	// 	if mCache[v] > 0 {
	// 		mCache[v] -= 1
	// 		res = append(res, v)
	// 	}
	// }

	//方案2 已排序
	pos1, pos2 := 0, 0
	sort.Ints(nums1)
	sort.Ints(nums2)

	//双指针写法
	for pos1 < len(nums1) && pos2 < len(nums2) {
		if nums1[pos1] == nums2[pos2] {
			res = append(res, nums1[pos1])
			pos1++
			pos2++
		} else if nums1[pos1] > nums2[pos2] {
			pos2++
		} else {
			pos1++
		}
	}

	// return []int{0}
	return res
}

// @lc code=end

