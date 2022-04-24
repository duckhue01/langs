package main

import (
	"reflect"
	"testing"
	"time"
)

const sleep = "sleep"
const write = "write"

type SpyCountdownOperation struct {
	Calls []string
}

func (s *SpyCountdownOperation) Sleep() {
	s.Calls = append(s.Calls, sleep)

}

func (s *SpyCountdownOperation) Write(p []byte) (int, error) {
	s.Calls = append(s.Calls, write)
	return 0, nil
}

func TestCountDown(t *testing.T) {
	t.Run("sleep before every print and prints 3 to Go", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperation{}
		Countdown(spySleepPrinter, spySleepPrinter)
		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfiguarableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}

	sleeper := ConfiguarableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
