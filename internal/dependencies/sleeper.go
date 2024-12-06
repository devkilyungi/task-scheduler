package dependencies

import "time"

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleepFn  func(time.Duration)
}

func NewConfigurableSleeper(duration time.Duration, sleepFn func(time.Duration)) *ConfigurableSleeper {
	return &ConfigurableSleeper{duration, sleepFn}
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleepFn(c.duration)
}
