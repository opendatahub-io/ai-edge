package support

import (
	"fmt"
	"time"
)

// WaitFor call a given function over a duration each tick given
// if the function return true then it is complete
// if the function ever returns an error then we stop checking and return
func WaitFor(duration time.Duration, tickRate time.Duration, f func() (bool, error)) error {
	timeout := time.After(duration)
	tick := time.Tick(tickRate)

	for {
		select {
		case <-timeout:
			return fmt.Errorf("waiting for operation to complete timedout after %v", duration.String())
		case <-tick:
			complete, err := f()
			if err != nil {
				return err
			}

			if complete {
				return nil
			}
		}
	}
}
