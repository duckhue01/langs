package main

import (
	"os"
	"time"
)

// type DefaultSleeper struct{}

// func (d *DefaultSleeper) Sleep() {
// 	time.Sleep(1 * time.Second)
// }

type ConfiguarableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfiguarableSleeper) Sleep() {
	c.sleep(c.duration)
}

func main() {
	Countdown(os.Stdout, &ConfiguarableSleeper{1 * time.Second, time.Sleep})
}
