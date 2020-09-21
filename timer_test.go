package timer

import (
	"testing"
	"time"
	"fmt"
)
func TestBackoffUntil(t *testing.T) {
	rt := NewRealTimer(3*time.Second)
	fmt.Println("begin")

	stopCh := make(chan struct{})
	go func() {
		// After 11 seconds, close the stopCh
		time.Sleep(11*time.Second)
		close(stopCh)
	}()

	rt.Until(func() {
		fmt.Println("hello")
	},stopCh)
	fmt.Println("end")
}
