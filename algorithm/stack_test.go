//go:build go1.18

package algorithm

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	n := int(1e3)
	samples := make([]int, n)
	for i := 0; i < n; i++ {
		samples[i] = rand.Intn(1e9)
	}

	st := NewStack[int]()
	assert.True(t, st.Empty())

	for i := 0; i < n; i++ {
		st.Push(samples[i])

		assert.False(t, st.Empty())
		assert.Equal(t, i+1, st.Size())
		assert.Equal(t, samples[i], st.Top())
	}

	for i := n - 1; i >= 0; i-- {
		assert.Equal(t, samples[i], st.Top())

		tmp := st.Pop()

		assert.Equal(t, i, st.Size())
		assert.Equal(t, samples[i], tmp)
	}

	assert.True(t, st.Empty())
}

func BenchmarkStack(b *testing.B) {
	loop := int(1e7)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		st := NewStack[int]()

		for i := 0; i < loop; i++ {
			st.Push(rand.Intn(1e9))
		}

		for i := 0; i < loop; i++ {
			st.Pop()
		}
	}

	b.ReportMetric(float64(b.N)/float64(b.Elapsed().Seconds()), "op/sec")
}
