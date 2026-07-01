//go:build linux

package gouuidv7

import "golang.org/x/sys/unix"

var hasSubMilliClock = false

func init() {
	var res unix.Timespec
	if err := unix.ClockGetres(unix.CLOCK_REALTIME, &res); err != nil {
		return
	}
	hasSubMilliClock = res.Sec == 0 && res.Nsec < 1e6
}
