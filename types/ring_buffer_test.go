package types_test

import (
	"testing"

	"github.com/flashbots/chain-monitor/types"
	"github.com/stretchr/testify/assert"
)

func TestRingBug(t *testing.T) {
	b := types.NewRingBuffer[int](42, 4)

	{
		assert.Equal(t, b.Length(), 0)
		_, nok := b.At(42)
		assert.False(t, nok)

		_, nok = b.Pop()
		assert.False(t, nok)
	}

	{
		b.Push(42)
		assert.Equal(t, b.Length(), 1)
		v, ok := b.At(42)
		assert.True(t, ok)
		assert.Equal(t, 42, v)

		_, nok := b.At(43)
		assert.False(t, nok)
	}

	{
		b.Push(43)
		b.Push(44)
		assert.Equal(t, b.Length(), 3)

		v, ok := b.At(42)
		assert.True(t, ok)
		assert.Equal(t, 42, v)

		v, ok = b.At(43)
		assert.True(t, ok)
		assert.Equal(t, 43, v)

		v, ok = b.At(44)
		assert.True(t, ok)
		assert.Equal(t, 44, v)

		_, nok := b.At(45)
		assert.False(t, nok)
	}

	{
		v, ok := b.Pop()
		assert.True(t, ok)
		assert.Equal(t, 42, v)

		assert.Equal(t, 2, b.Length())

		_, nok := b.At(42)
		assert.False(t, nok)

		v, ok = b.At(43)
		assert.True(t, ok)
		assert.Equal(t, 43, v)

		v, ok = b.At(44)
		assert.True(t, ok)
		assert.Equal(t, 44, v)
	}

	{
		b.Push(45)

		assert.Equal(t, 3, b.Length())

		v, ok := b.Pop()
		assert.True(t, ok)
		assert.Equal(t, 43, v)

		assert.Equal(t, 2, b.Length())

		_, nok := b.At(42)
		assert.False(t, nok)

		_, nok = b.At(43)
		assert.False(t, nok)

		v, ok = b.At(44)
		assert.True(t, ok)
		assert.Equal(t, 44, v)

		v, ok = b.At(45)
		assert.True(t, ok)
		assert.Equal(t, 45, v)
	}

	{
		b.Push(46)

		v, ok := b.Pop()
		assert.True(t, ok)
		assert.Equal(t, 44, v)

		assert.Equal(t, 2, b.Length())

		_, nok := b.At(42)
		assert.False(t, nok)

		_, nok = b.At(43)
		assert.False(t, nok)

		_, nok = b.At(44)
		assert.False(t, nok)

		v, ok = b.At(45)
		assert.True(t, ok)
		assert.Equal(t, 45, v)

		v, ok = b.At(46)
		assert.True(t, ok)
		assert.Equal(t, 46, v)
	}
}
