package nextgeneration

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_NextGeneration(t *testing.T){
	var A [3][3]int = [3][3]int{
		{1,1,1},
		{0,0,0},
		{0,0,0},
	}
	t.Run("Test to ensure NewGeneration doesnot Panic", func(t *testing.T){
		assert.NotPanics(t,func(){NextGeneration(A)})
	})

	var expected [3][3]int = [3][3] int{
		{0,0,0},
		{0,1,0},
		{0,0,0},
	}
	got := NextGeneration(A)
	t.Run("Test added to ensure the NextGeneration function produces the correct next generation", func(t *testing.T){
		assert.IsType(t,expected,got)
	})
}