package asteroid

import "time"

type Timer struct {
	clock    Clock
	duration time.Duration
}

func NewTimer(duration time.Duration) *Timer {
	return &Timer{duration: duration}
}

func (t *Timer) SetDuration(duration time.Duration) {
	t.duration = duration
}

func (t *Timer) GetDuration() time.Duration {
	return t.duration
}

func (t *Timer) Start() *Timer {
	t.clock.SetStartTime(time.Now())
	return t
}

func (t *Timer) IsFinished() bool {
	if t.clock.GetElapsedTime() >= t.duration {
		t.clock.Restart()
		return true
	}
	return false
}
