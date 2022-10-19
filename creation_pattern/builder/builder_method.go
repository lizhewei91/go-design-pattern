package builder

import "fmt"

const (
	defaultName     = "test"
	defaultMaxTotal = 50
	defaultMaxIdle  = 15
	defaultMinIdle  = 3
)

type ResourcePoolConfig struct {
	name     string // len() <100
	maxTotal int    // (10,100)
	maxIdle  int    // (10,20)
	minIdle  int    // (0,5)
}

type ResourcePoolConfigBuilder struct {
	input
}

type input struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
	allErr   []error
}

func NewResourcePoolConfigBuilder() *ResourcePoolConfigBuilder {
	return &ResourcePoolConfigBuilder{}
}

func (b *ResourcePoolConfigBuilder) WithName(name string) *ResourcePoolConfigBuilder {
	if len(name) > 100 {
		b.input.allErr = append(b.input.allErr, fmt.Errorf("name length must <100"))
	}
	b.input.name = name
	return b
}

func (b *ResourcePoolConfigBuilder) WithMaxTotal(maxTotal int) *ResourcePoolConfigBuilder {
	if maxTotal < 10 || maxTotal > 100 {
		b.input.allErr = append(b.input.allErr, fmt.Errorf("maxtotal must 10< maxTotal <100\n"))
	}
	b.input.maxTotal = maxTotal
	return b
}

func (b *ResourcePoolConfigBuilder) WithMinIdle(minIdle int) *ResourcePoolConfigBuilder {
	if minIdle > 5 || minIdle < 0 {
		b.input.allErr = append(b.input.allErr, fmt.Errorf("minIdle must 0< minIdle <5\n"))
	}
	b.minIdle = minIdle
	return b
}

func (b *ResourcePoolConfigBuilder) WithMaxIdle(maxIdle int) *ResourcePoolConfigBuilder {
	if maxIdle < 10 || maxIdle > 20 {
		b.input.allErr = append(b.input.allErr, fmt.Errorf("maxIdle must 10< maxIdle <20\n"))
	}
	b.maxIdle = maxIdle
	return b
}

func (b *ResourcePoolConfigBuilder) Builder() (*ResourcePoolConfig, []error) {
	if b.name == "" {
		b.name = defaultName
	}
	if b.maxTotal == 0 {
		b.maxTotal = defaultMaxTotal
	}
	if b.maxIdle == 0 {
		b.maxIdle = defaultMaxIdle
	}
	if b.minIdle == 0 {
		b.minIdle = defaultMinIdle
	}

	if len(b.input.allErr) != 0 {
		return nil, b.input.allErr
	}

	return &ResourcePoolConfig{
		name:     b.name,
		maxTotal: b.maxTotal,
		maxIdle:  b.maxIdle,
		minIdle:  b.minIdle,
	}, nil
}

func NewResourcePoolConfig() (*ResourcePoolConfig, []error) {
	rpf, allErr := NewResourcePoolConfigBuilder().
		WithName("lzw").
		WithMaxTotal(55).
		WithMaxIdle(12).
		WithMinIdle(3).
		Builder()
	if len(allErr) != 0 {
		return nil, allErr
	}
	return rpf, nil
}
