package timer
import "time"

func NewRealTimer(resetTime time.Duration)(*realTimer) {
	return &realTimer{
		timer:time.NewTimer(resetTime),
		resetTime:resetTime,
	}
}

func (t *realTimer)Until(f func(), stopCh <-chan struct{}) {
	defer func() {
		t.timer.Stop()
	}()

	for {
		select {
		case <-stopCh:
			return
		default:
		}

		t.Reset(t.resetTime)
		f()
		select {
		case <-stopCh:
			return
		case <-t.C():
		}
	}
}


type realTimer struct {
	timer *time.Timer
	resetTime time.Duration
}

// C returns the underlying timer's channel.
func (r *realTimer) C() <-chan time.Time {
	return r.timer.C
}

// Reset calls Reset() on the underlying timer.
func (r *realTimer) Reset(d time.Duration) bool {
	return r.timer.Reset(d)
}