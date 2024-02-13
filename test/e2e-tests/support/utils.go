package support

import (
	"fmt"
	"time"
)

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
