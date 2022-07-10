package two_sum

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTwoSumBasic(t *testing.T) {
	target := 23
	subArr := []int{7, 8, 11, 13, 9, 3, 7, 2}
	b := twoSum(target, subArr)
	require.True(t, b)

	b = TwoSumTable(target, subArr)
	require.True(t, b)
}

func TestTwoSumBasic2(t *testing.T) {
	target := 23
	subArr := []int{7, 9, 13}
	b := twoSum(target, subArr)
	require.True(t, b)

	b = TwoSumTable(target, subArr)
	require.True(t, b)
}

func TestTwoSumBasic3(t *testing.T) {
	target := 24
	subArr := []int{7, 9, 13}
	b := twoSum(target, subArr)
	require.False(t, b)

	b = TwoSumTable(target, subArr)
	require.False(t, b)
}
