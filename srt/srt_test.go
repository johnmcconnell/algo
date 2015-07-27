package srt

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	assert := assert.New(t)

	ls := []int{-1, 1, -1, 1, -1, 1}
	sLs := []int{-1, -1, -1, 1, 1, 1}

	MergeSort(ls)
	assert.Equal(sLs, ls, "Sorted list matches after sorting")

	ls = []int{-1, 1, -1, 1, -1, 1, -1}
	sLs = []int{-1, -1, -1, -1, 1, 1, 1}

	MergeSort(ls)
	assert.Equal(sLs, ls, "Sorted list matches after sorting")

	ls = []int{
		-346,
		62393,
		4,
		3485,
		0,
		0,
		-852823,
		-2835,
		2,
		-38,
		2,
		203,
		-1,
		2093,
	}

	sLs = []int{
		-852823,
		-2835,
		-346,
		-38,
		-1,
		0,
		0,
		2,
		2,
		4,
		203,
		2093,
		3485,
		62393,
	}

	MergeSort(ls)
	assert.Equal(sLs, ls, "Sorted list matches after sorting")
}
