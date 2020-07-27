package main

import (
	"testing"
	"sync")

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times", func (t *testing.T){
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		want := 3
		if counter.Value() != 3 {
			t.Errorf("got %d, want %d", counter.Value(), want)
		}
	})
	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()
		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i<wantedCount;i++{
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			} (&wg)
		}
		wg.Wait()
		assertCounter(t, counter, wantedCount)
	})
}
func assertCounter(t *testing.T, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
