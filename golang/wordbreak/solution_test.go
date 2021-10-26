package wordbreak

import (
	"golang/helper"
	"testing"
)

func Test(t *testing.T) {
	s, wordDict := "leetcode", []string{"leet", "code"}
	helper.Assert(wordBreak(s, wordDict), true, t)

	s, wordDict = "applepenapple", []string{"apple", "pen"}
	helper.Assert(wordBreak(s, wordDict), true, t)

	s, wordDict = "catsandog", []string{"cats", "dog", "sand", "and", "cat"}
	helper.Assert(wordBreak(s, wordDict), false, t)

	s = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	wordDict = []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	helper.Assert(wordBreak(s, wordDict), false, t)

	s = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	wordDict = []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	helper.Assert(wordBreak(s, wordDict), false, t)
}
