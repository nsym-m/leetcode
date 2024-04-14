package main

/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 *
 * https://leetcode.com/problems/product-of-array-except-self/description/
 *
 * algorithms
 * Medium (65.11%)
 * Likes:    22063
 * Dislikes: 1353
 * Total Accepted:    2.5M
 * Total Submissions: 3.8M
 * Testcase Example:  '[1,2,3,4]'
 *
 * Given an integer array nums, return an array answer such that answer[i] is
 * equal to the product of all the elements of nums except nums[i].
 *
 * The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit
 * integer.
 *
 * You must write an algorithm that runs in O(n) time and without using the
 * division operation.
 *
 *
 * Example 1:
 * Input: nums = [1,2,3,4]
 * Output: [24,12,8,6]
 * Example 2:
 * Input: nums = [-1,1,0,-3,3]
 * Output: [0,0,9,0,0]
 *
 *
 * Constraints:
 *
 *
 * 2 <= nums.length <= 10^5
 * -30 <= nums[i] <= 30
 * The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit
 * integer.
 *
 *
 *
 * Follow up: Can you solve the problem in O(1) extra space complexity? (The
 * output array does not count as extra space for space complexity analysis.)
 *
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	prefix := 1
	for i := 0; i < n; i++ {
		ans[i] = prefix
		prefix *= nums[i]
	}

	suffix := 1
	for i := n - 1; i >= 0; i-- {
		ans[i] *= suffix
		suffix *= nums[i]
	}

	return ans

	/*
	   このアルゴリズムは、２つのfor文を使用して、配列の各インデックスにおいて、自身を除く他の全要素の積を計算します。この方法は、インデックスごとに左側の要素の積と右側の要素の積を個別に計算し、最後にそれらを掛け合わせることで実現されます。

	   ## 最初のfor文 (for i := 0; i < n; i++)
	   このループでは、左側の要素の累積積を計算し、ans 配列に格納します。ループが進むにつれて、curr は nums 配列の現在の要素 nums[i] に更新される前に、ans[i] に乗算されます。これにより、ans[i] はインデックス i までの全ての要素の積を持つことになりますが、nums[i] 自体は含まれません（curr が nums[i] を乗算するのは ans[i] に値を設定した後です）。

	   例えば、nums = [1, 2, 3, 4] の場合、最初のループの後で ans は次のようになります：

	   ans[0] = 1 (何も掛けられていない)
	   ans[1] = 1 (1のみが掛けられる)
	   ans[2] = 2 (1と2の積)
	   ans[3] = 6 (1、2、3の積)

	   ## 二番目のfor文 (for i := n - 1; i >= 0; i--)
	   このループでは、右側の要素の累積積を計算します。ここでも curr を使用して、ans[i] に対して右から左への積を掛け合わせていきます。このステップにより、ans[i] には以前のループで計算された左側の要素の積と、このループで計算される右側の要素の積が乗算されます。この結果、ans[i] は nums[i] を除く全要素の積が得られます。

	   上記の例で続けると、二番目のループの後で ans は次のようになります：

	   ans[0] = 24 (2、3、4の積)
	   ans[1] = 12 (1、3、4の積)
	   ans[2] = 8 (1、2、4の積)
	   ans[3] = 6 (1、2、3の積)
	   このようにして、２つのループを通じて、配列の各要素に対して自身を除いた他の要素の積が ans 配列に格納されるのです。
	*/
}

// @lc code=end
