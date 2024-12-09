package dependencies

import (
	"testing"
	"time"
)

type SpySleeper struct {
	durationSlept time.Duration
}

func (s *SpySleeper) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 3 * time.Second
	spySleeper := &SpySleeper{}
	sleeper := ConfigurableSleeper{sleepTime, spySleeper.Sleep}

	sleeper.Sleep()

	if spySleeper.durationSlept != sleepTime {
		t.Fatalf("expected to sleep for %v, but slept for %v", sleepTime, spySleeper.durationSlept)
	}
}
