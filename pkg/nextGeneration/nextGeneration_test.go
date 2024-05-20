package nextgeneration

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_NextGeneration(t *testing.T){
	var A [3][3]int = [3][3]int{
		{1,2,3},
		{1,2,3},
		{1,2,3},
	}
	t.Run("Test to ensure NewGeneration doesnot Panic", func(t *testing.T){
		assert.NotPanics(t,func(){NextGeneration(A)})
	})
}