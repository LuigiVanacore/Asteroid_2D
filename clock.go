package asteroid

import "time"

type Clock struct {
	startTime time.Time
}

func (c *Clock) Restart() time.Duration {
	elapsedTime := time.Now().Sub(c.startTime)
	c.startTime = time.Now()
	return elapsedTime
}

func (c *Clock) GetElapsedTime() time.Duration {
	return time.Now().Sub(c.startTime)
}

func (c *Clock) SetStartTime(startTime time.Time) {
	c.startTime = startTime
}
