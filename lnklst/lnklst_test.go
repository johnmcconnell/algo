package lnklst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	assert := assert.New(t)

	lst := NewSingleLinkedList()

	a := "a"
	b := "b"
	c := "c"
	ws := []*string{&a, &b, &c}

	for _, w := range ws {
		lst.Add(w)
	}

	for i := range ws {
		assert.Equal(ws[i], lst.Get(i), "List indexes match slices")
	}
}

func TestDel(t *testing.T) {
	assert := assert.New(t)

	lst := NewSingleLinkedList()

	a := "a"
	b := "b"
	c := "c"
	ws := []*string{&a, &b, &c}

	for _, w := range ws {
		lst.Add(w)
	}

	d := lst.Del(1)

	assert.Equal(ws[0], lst.Get(0), "List indexes match slices")
	assert.Equal(ws[1], d, "List indexes match slices")
	assert.Equal(ws[2], lst.Get(1), "List indexes match slices")
}
