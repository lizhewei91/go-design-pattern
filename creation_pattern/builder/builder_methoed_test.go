package builder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewResourcePoolConfig(t *testing.T) {
	got, allErr := NewResourcePoolConfigBuilder().
		WithName("lzw").
		WithMaxTotal(55).
		WithMaxIdle(12).
		WithMinIdle(3).
		Builder()

	want := &ResourcePoolConfig{
		name:     "lzw",
		maxTotal: 55,
		maxIdle:  12,
		minIdle:  3,
	}
	assert.Equal(t, 0, len(allErr))
	assert.Equal(t, want, got)
}
