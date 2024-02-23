package main

import "fmt"

/*
 * @lc app=leetcode id=1768 lang=golang
 *
 * [1768] Merge Strings Alternately
 *
 * https://leetcode.com/problems/merge-strings-alternately/description/
 *
 * algorithms
 * Easy (82.42%)
 * Likes:    3546
 * Dislikes: 76
 * Total Accepted:    660.3K
 * Total Submissions: 831.4K
 * Testcase Example:  '"abc"\n"pqr"'
 *
 * You are given two strings word1 and word2. Merge the strings by adding
 * letters in alternating order, starting with word1. If a string is longer
 * than the other, append the additional letters onto the end of the merged
 * string.
 *
 * Return the merged string.
 *
 *
 * Example 1:
 *
 *
 * Input: word1 = "abc", word2 = "pqr"
 * Output: "apbqcr"
 * Explanation: The merged string will be merged as so:
 * word1:  a   b   c
 * word2:    p   q   r
 * merged: a p b q c r
 *
 *
 * Example 2:
 *
 *
 * Input: word1 = "ab", word2 = "pqrs"
 * Output: "apbqrs"
 * Explanation: Notice that as word2 is longer, "rs" is appended to the end.
 * word1:  a   b
 * word2:    p   q   r   s
 * merged: a p b q   r   s
 *
 *
 * Example 3:
 *
 *
 * Input: word1 = "abcd", word2 = "pq"
 * Output: "apbqcd"
 * Explanation: Notice that as word1 is longer, "cd" is appended to the end.
 * word1:  a   b   c   d
 * word2:    p   q
 * merged: a p b q c   d
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= word1.length, word2.length <= 100
 * word1 and word2 consist of lowercase English letters.
 *
 */

// @lc code=start
func mergeAlternately(word1 string, word2 string) string {
	w1len := len(word1)
	w2len := len(word2)
	l := w1len
	if w1len < w2len {
		l = w2len
	}
	s := ""

	for i := 0; i < l; i++ {
		if w1len > i && w2len > i {
			s += string(word1[i]) + string(word2[i])
		}
		if w1len <= i {
			s += string(word2[i:])
			break
		}
		if w2len <= i {
			s += string(word1[i:])
			break
		}
	}
	return s
}

// @lc code=end

func main() {
	fmt.Println(mergeAlternately("cf", "eee"))
}
