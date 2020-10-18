package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_chop(t *testing.T) {
	t.Run("testing chop function ", func(t *testing.T) {
		assert.Equal(t, -1, chop(3, []int{}))
		assert.Equal(t, -1, chop(3, []int{1}))
		assert.Equal(t, 0, chop(1, []int{1}))
		assert.Equal(t, 0, chop(1, []int{1, 3, 5}))
		assert.Equal(t, 1, chop(3, []int{1, 3, 5}))
		assert.Equal(t, 2, chop(5, []int{1, 3, 5}))
		assert.Equal(t, -1, chop(0, []int{1, 3, 5}))
		assert.Equal(t, -1, chop(2, []int{1, 3, 5}))
		assert.Equal(t, -1, chop(4, []int{1, 3, 5}))
		assert.Equal(t, -1, chop(6, []int{1, 3, 5}))
		assert.Equal(t, 0, chop(1, []int{1, 3, 5, 7}))
		assert.Equal(t, 1, chop(3, []int{1, 3, 5, 7}))
		assert.Equal(t, 2, chop(5, []int{1, 3, 5, 7}))
		assert.Equal(t, 3, chop(7, []int{1, 3, 5, 7}))
		assert.Equal(t, -1, chop(0, []int{1, 3, 5, 7}))
		assert.Equal(t, -1, chop(2, []int{1, 3, 5, 7}))
		assert.Equal(t, -1, chop(4, []int{1, 3, 5, 7}))
		assert.Equal(t, -1, chop(6, []int{1, 3, 5, 7}))
		assert.Equal(t, -1, chop(8, []int{1, 3, 5, 7}))
	})
}
