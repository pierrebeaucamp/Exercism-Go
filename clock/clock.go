package clock

import "fmt"

type clock struct {
	hour, minute int
}

func New(hour, minute int) clock {
	return clock{hour, minute}
}

func (c clock) Add(minute int) clock {
	var m int = (c.hour * 60) + c.minute + minute

	m %= 24 * 60
	if m < 0 {
		m += 24 * 60
	}

	c.hour = m / 60
	c.minute = m % 60
	return c
}

func (c clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
