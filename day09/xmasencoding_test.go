package main

import (
	. "./encoding"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestFindFirstFaulty(t *testing.T) {
	nums := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}

	faultyPos := FindFirstFaulty(nums, 5)

	assert.Equal(t, nums[faultyPos], 127)
}

func TestFindSet(t *testing.T) {
	nums := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182}

	set := FindContiguousSet(nums, 127)

	assert.Equal(t, set, []int{15, 25, 47, 40})
}
