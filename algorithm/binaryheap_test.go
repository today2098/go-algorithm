//go:build go1.18

package algorithm

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryHeap(t *testing.T) {
	n := int(1e3)
	samples := make([]int, n)
	for i := 0; i < n; i++ {
		samples[i] = rand.Intn(1e9)
	}

	pque := NewDefaultBinaryHeap[int]()
	assert.True(t, pque.Empty())

	mx := -1
	for i := 0; i < n; i++ {
		pque.Push(samples[i])
		mx = max(mx, samples[i])

		assert.False(t, pque.Empty())
		assert.Equal(t, i+1, pque.Size())
		assert.Equal(t, mx, pque.Top())
	}

	sort.Slice(samples, func(i, j int) bool {
		return samples[i] > samples[j]
	})

	for i := 0; i < n; i++ {
		tmp := pque.Pop()

		assert.Equal(t, n-1-i, pque.Size())
		assert.Equal(t, samples[i], tmp)
	}

	assert.True(t, pque.Empty())
}

func BenchmarkBinaryHeap(b *testing.B) {
	loop := int(1e7)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pque := NewDefaultBinaryHeap[int]()

		for i := 0; i < loop; i++ {
			pque.Push(rand.Intn(1e9))
		}

		for i := 0; i < loop; i++ {
			pque.Pop()
		}
	}

	b.ReportMetric(float64(b.N)/float64(b.Elapsed().Seconds()), "op/sec")
}
