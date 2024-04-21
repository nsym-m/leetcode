package main

/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 *
 * https://leetcode.com/problems/move-zeroes/description/
 *
 * algorithms
 * Easy (61.41%)
 * Likes:    16371
 * Dislikes: 435
 * Total Accepted:    2.9M
 * Total Submissions: 4.7M
 * Testcase Example:  '[0,1,0,3,12]'
 *
 * Given an integer array nums, move all 0's to the end of it while maintaining
 * the relative order of the non-zero elements.
 *
 * Note that you must do this in-place without making a copy of the array.
 *
 *
 * Example 1:
 * Input: nums = [0,1,0,3,12]
 * Output: [1,3,12,0,0]
 * Example 2:
 * Input: nums = [0]
 * Output: [0]
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 10^4
 * -2^31 <= nums[i] <= 2^31 - 1
 *
 *
 *
 * Follow up: Could you minimize the total number of operations done?
 */

// @lc code=start
func moveZeroes(nums []int) {

	zero := 0
	for i, v := range nums {
		if v == 0 {
			zero++
		} else {
			nums[i-zero] = v
		}
	}
	l := len(nums)
	for i := 0; i < zero; i++ {
		nums[l-1-i] = 0
	}
}

// @lc code=end
