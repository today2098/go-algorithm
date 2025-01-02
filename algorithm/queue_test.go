//go:build go1.18

package algorithm

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQueue(t *testing.T) {
	n := int(1e3)
	samples := make([]int, n)
	for i := 0; i < n; i++ {
		samples[i] = rand.Intn(1e9)
	}

	que := NewQueue[int]()
	assert.True(t, que.Empty())

	for i := 0; i < n; i++ {
		que.Push(samples[i])

		assert.False(t, que.Empty())
		assert.Equal(t, i+1, que.Size())
		assert.Equal(t, samples[0], que.Front())
	}

	for i := 0; i < n; i++ {
		assert.Equal(t, samples[i], que.Front())

		tmp := que.Pop()

		assert.Equal(t, n-1-i, que.Size())
		assert.Equal(t, samples[i], tmp)
	}

	assert.True(t, que.Empty())
}

func TestQueue_Panic(t *testing.T) {
	tests := []struct {
		name    string
		f       func(t *testing.T)
		wantErr any
	}{
		{
			name: "Front",
			f: func(t *testing.T) {
				que := NewQueue[int]()
				require.True(t, que.Empty())
				_ = que.Front()
			},
			wantErr: "Queue: empty",
		},
		{
			name: "Pop",
			f: func(t *testing.T) {
				que := NewQueue[int]()
				require.True(t, que.Empty())
				que.Pop()
			},
			wantErr: "Queue: empty",
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

func BenchmarkQueue(b *testing.B) {
	loop := int(1e7)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		que := NewQueue[int]()

		for i := 0; i < loop; i++ {
			que.Push(rand.Intn(1e9))
		}

		for i := 0; i < loop; i++ {
			que.Pop()
		}
	}

	b.ReportMetric(float64(b.N)/float64(b.Elapsed().Seconds()), "op/sec")
}
