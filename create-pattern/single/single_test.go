package single

import (
	"sync"
	"testing"
)

var wg sync.WaitGroup

func TestSingle(t *testing.T) {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			single := GetSingleInstance()
			t.Log(single)
		}()
	}
	wg.Wait()
}
