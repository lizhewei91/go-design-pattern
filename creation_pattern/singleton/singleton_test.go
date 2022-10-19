package singleton

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstance(t *testing.T) {
	ins1 := GetInstance()
	ins2 := GetInstance()
	assert.Equal(t, ins1, ins2)
}

func TestParallelSingleton(t *testing.T) {
	wg := sync.WaitGroup{}
	instances := [100]*singleton{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			ins := GetInstance()
			instances[n] = ins
		}(i)
	}

	wg.Wait()

	for i := 1; i < 100; i++ {
		assert.Equal(t, instances[i], instances[i-1])
	}
}
