package clock

import (
	"fmt"
)

type Clock struct {                     // define the clock type
	h, m int
}

func AbsMinutes(x int) int {
	if x < 0 {
		return 60+x
	}
	return x
}

func AbsHours(x int) int {
	if x < 0 {
		return 24+x
	}
	return x
}

func New(hours, minutes int) Clock {	// a "constructor"
	var clock Clock

	_minutes := 60 * hours + minutes
	_minutes = _minutes % (60 * 24)
	if _minutes < 0 {
		_minutes = 60 * 24 + _minutes
	}

	clock.m = _minutes % 60
	clock.h = _minutes / 60

	return clock 
}

func (c Clock) String() string {	// a "stringer"
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

func (c Clock) Add(minutes int) Clock {
	var clock Clock

	_minutes := 60 * c.h + c.m + minutes
	_minutes = _minutes % (60 * 24)
	if _minutes < 0 {
		_minutes = 60 * 24 + _minutes
	}

	clock.m = _minutes % 60
	clock.h = _minutes / 60

	return clock 
}

func (c Clock) Subtract(minutes int) Clock {
	return c.Add(-minutes)
}
