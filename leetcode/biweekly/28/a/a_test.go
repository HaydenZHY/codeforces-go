// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [a]")
	examples := [][]string{
		{
			`[8,4,6,2,3]`, 
			`[4,2,4,2,3]`,
		},
		{
			`[1,2,3,4,5]`, 
			`[1,2,3,4,5]`,
		},
		{
			`[10,1,1,6]`, 
			`[9,0,1,6]`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, finalPrices, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-28/problems/final-prices-with-a-special-discount-in-a-shop/
