package singleton

import (
	"reflect"
	"sync"
	"testing"
)

const parCount = 100

func TestSingleton(t *testing.T) {
	ins1 := GetInstance("lzw", 18)
	ins2 := GetInstance("lzw", 18)
	if !reflect.DeepEqual(ins1, ins2) {
		t.Errorf("instance is not equal")
	}
}

func TestParallelSingleton(t *testing.T) {
	start := make(chan struct{})
	var (
		wg = sync.WaitGroup{}
		mu = sync.Mutex{}
	)

	wg.Add(parCount)
	singletons := make(map[int]Singleton, parCount)
	for i := 0; i < parCount; i++ {
		go func(index int) {
			<-start
			mu.Lock()
			singletons[index] = GetInstance("lzw", 18)
			mu.Unlock()
			wg.Done()
		}(i)
	}
	close(start)
	wg.Wait()
	for i := 1; i < parCount; i++ {
		if !reflect.DeepEqual(singletons[i], singletons[i-1]) {
			t.Errorf("singleton %v is not equal singleton %v\n", i, i-1)
		}
	}
}
