/*
 * @Author: Y.t
 * @Date: 2022-06-12 22:13:42
 * @LastEditors: Y.t
 * @LastEditTime: 2022-06-13 22:49:38
 * @Description:
 */
/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}

	var prefix = strs[0]
	var length = len(prefix)
	for _, v := range strs {

		//缩减前缀, 直到匹配
		for strings.Index(v, prefix[:length]) != 0 {
			if length < 1 {
				return ""
			}
			length--
		}
	}

	return prefix[:length]
}

// @lc code=end

