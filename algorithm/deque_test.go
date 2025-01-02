//go:build go1.18

package algorithm

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeque(t *testing.T) {
	n := int(1e3)
	samples := make([]int, n)
	for i := 0; i < n; i++ {
		samples[i] = rand.Intn(1e9)
	}

	dque := NewDeque[int]()
	assert.True(t, dque.Empty())

	l, r := 0, n
	for l < r {
		if rand.Intn(2) == 1 {
			r--
			dque.PushBack(samples[r])
		} else {
			dque.PushFront(samples[l])
			l++
		}

		assert.Equal(t, n-(r-l), dque.Size())
		assert.Equal(t, samples[(l-1+n)%n], dque.Front())
		assert.Equal(t, samples[r%n], dque.Back())
	}

	for r-l < n {
		assert.Equal(t, samples[(l-1+n)%n], dque.Front())
		assert.Equal(t, samples[r%n], dque.Back())

		if rand.Intn(2) == 1 {
			assert.Equal(t, samples[r%n], dque.PopBack())
			r++
		} else {
			l--
			assert.Equal(t, samples[(l+n)%n], dque.PopFront())
		}

		assert.Equal(t, n-(r-l), dque.Size())
	}

	assert.True(t, dque.Empty())
}

func TestDeque_Panic(t *testing.T) {
	tests := []struct {
		name    string
		f       func(t *testing.T)
		wantErr any
	}{
		{
			name: "Front",
			f: func(t *testing.T) {
				dque := NewDeque[int]()
				require.True(t, dque.Empty())
				_ = dque.Front()
			},
			wantErr: "Deque: empty",
		},
		{
			name: "Back",
			f: func(t *testing.T) {
				dque := NewDeque[int]()
				require.True(t, dque.Empty())
				_ = dque.Back()
			},
			wantErr: "Deque: empty",
		},
		{
			name: "PopFront",
			f: func(t *testing.T) {
				dque := NewDeque[int]()
				require.True(t, dque.Empty())
				_ = dque.PopFront()
			},
			wantErr: "Deque: empty",
		},
		{
			name: "PopBack",
			f: func(t *testing.T) {
				dque := NewDeque[int]()
				require.True(t, dque.Empty())
				_ = dque.PopBack()
			},
			wantErr: "Deque: empty",
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

func BenchmarkDeque(b *testing.B) {
	loop := int(1e7)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		dque := NewDeque[int]()

		for i := 0; i < loop; i++ {
			if rand.Intn(2) == 1 {
				dque.PushBack(rand.Intn(1e9))
			} else {
				dque.PushFront(rand.Intn(1e9))
			}
		}

		for i := 0; i < loop; i++ {
			if rand.Intn(2) == 1 {
				dque.PopBack()
			} else {
				dque.PopFront()
			}
		}
	}

	b.ReportMetric(float64(b.N)/float64(b.Elapsed().Seconds()), "op/sec")
}
