package gigasecond

import "time"

const Gigasecond = 1e18

func AddGigasecond(input time.Time) time.Time {
	return input.Add(Gigasecond)
}