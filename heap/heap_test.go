package heap

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTreeHeap(t *testing.T) {
	assert := assert.New(t)

	XS := []int{4, 5, 1, 3, 4, 5, 9, 5}
	OS := []int{1, 3, 4, 4, 5, 5, 5, 9}

	h := Array()

	for _, x := range XS {
		h.Add(x)
	}

	for _, o := range OS {
		v := h.Pop()

		assert.Equal(o, v, "The popped value is in the same order as sorted array")
	}
}

func TestMultipleOrderHeap(t *testing.T) {
	assert := assert.New(t)

	XS := []int{4, 5, 1, 3, 4, 5, 9, 5}
	OS := []int{1, 3, 4, 4, 5, 5, 5, 9}

	h := Array()

	for _, x := range XS {
		h.Add(x)
	}

	v0 := h.Pop()
	v1 := h.Pop()
	v2 := h.Pop()

	h.Add(v1)
	h.Add(v2)

	v3 := h.Pop()

	h.Add(v0)
	h.Add(v3)
	// Balance even

	v0 = h.Pop()
	h.Add(v0)

	v1 = h.Pop()
	v2 = h.Pop()

	h.Add(v2)
	h.Add(v1)

	for _, o := range OS {
		v := h.Pop()

		assert.Equal(o, v, "The popped value is in the same order as sorted array")
	}
}
