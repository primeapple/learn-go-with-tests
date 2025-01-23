package counter

import (
	"sync"
	"testing"
)

func assertCounterValue(t testing.TB, counter *Counter, want int) {
	t.Helper()

	got := counter.Value()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounterValue(t, counter, 3)
	})

	t.Run("it runs safely asynchronously", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < 1000; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
        wg.Wait()

        assertCounterValue(t, counter, 1000)
	})
}
