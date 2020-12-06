package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveSortedDupElements(t *testing.T)  {
	dups := []int{1,1,2,3,3,4,5,6,7}
	realLen := RemoveSortedDupElements(dups)
	assert.Equal(t,realLen,7)
}
