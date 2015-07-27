package lnklst

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert := assert.New(t)

	lst := NewSingleLinkedList()

	a := "a"
	b := "b"
	c := "c"
	ws := []*string{&a, &b, &c}

	for _, w := range ws {
		lst.Append(w)
	}

	for i := range ws {
		assert.Equal(ws[i], lst.Get(i), "List indexes match slices")
	}

}
