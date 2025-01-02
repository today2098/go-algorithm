//go:build go1.18

package algorithm

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestBinaryHeap_Panic(t *testing.T) {
	tests := []struct {
		name    string
		f       func(t *testing.T)
		wantErr any
	}{
		{
			name: "Top",
			f: func(t *testing.T) {
				pque := NewDefaultBinaryHeap[int]()
				require.True(t, pque.Empty())
				_ = pque.Top()
			},
			wantErr: "BinaryHeap: empty",
		},
		{
			name: "Pop",
			f: func(t *testing.T) {
				pque := NewDefaultBinaryHeap[int]()
				require.True(t, pque.Empty())
				pque.Pop()
			},
			wantErr: "BinaryHeap: empty",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				err := recover()
				assert.Equal(t, test.wantErr, err)
			}()

			test.f(t)
		})
	}
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
