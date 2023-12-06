// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1907/problem/G
// https://codeforces.com/problemset/status/1907/problem/G
func Test_cf1907G(t *testing.T) {
	testCases := [][2]string{
		{
			`8
5
11101
4 3 4 2 2
2
10
2 1
10
0000000011
9 10 10 7 10 9 9 9 10 2
10
1000111101
9 3 8 9 2 1 3 7 2 7
10
0001101010
5 7 6 10 8 3 6 6 2 2
10
0101100010
8 7 7 9 9 4 1 4 2 7
10
1010111010
7 9 10 7 7 2 8 6 10 4
10
1110000001
3 10 10 1 10 8 6 3 2 1`,
			`3
1 5 3 
-1
1
9 
5
5 6 10 2 3 
6
4 9 5 10 8 7 
3
5 4 9 
6
1 3 5 9 7 8 
2
2 1`,
		}, {
			`1
10
0010011100
6 7 9 1 3 3 1 3 8 7`,
			`3
7 1 8 `,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1907G)
}
