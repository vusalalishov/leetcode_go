package p1_test

import (
	"github.com/vusalalishov/leetcode_go/easy/p1"
	"testing"
)

func TestTwoSum(t *testing.T) {
	sum := p1.TwoSum([]int{2, 7, 11, 15}, 9)
	if len(sum) != 2 && sum[0] != 0 && sum[1] != 1 {
		t.Fail()
	}
}
