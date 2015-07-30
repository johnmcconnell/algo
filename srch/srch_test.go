package srch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	SLs = []int{
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
)

func TestSingleBinarySearch(t *testing.T) {
	assert := assert.New(t)
	idx := SingleBinarySearch(
		len(SLs),
		func(x int) int {
			return SLs[x] - 4
		},
	)

	assert.Equal(
		SLs[idx],
		4,
		"Found the number in the list",
	)

	idx = SingleBinarySearch(
		len(SLs),
		func(x int) int {
			return SLs[x] - 44
		},
	)

	assert.Equal(
		idx,
		-1,
		"Index of item not found",
	)
}

func TestMultBinarySearch(t *testing.T) {
	assert := assert.New(t)

	var items []int

	MultBinarySearch(
		len(SLs),
		func(x int) int {
			return SLs[x] - 2
		},
		func(x int) {
			items = append(items, SLs[x])
		},
	)

	assert.Equal(
		len(items),
		2,
		"Found 2 items",
	)

	for _, i := range items {
		assert.Equal(
			i,
			2,
			"Index maps to item found",
		)
	}

	items = []int{}

	MultBinarySearch(
		len(SLs),
		func(x int) int {
			return SLs[x] - 9378121
		},
		func(x int) {
			items = append(items, SLs[x])
		},
	)

	assert.Equal(
		len(items),
		0,
		"Found 0 ids",
	)
}
