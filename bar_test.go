package bar

import (
	"testing"
	"time"
)

func TestBar(t *testing.T) {
	bar := NewBar().WithGraph("█").NewOption(0, 100)
	for i := 0; i <= 100; i++ {
		bar.Play(int64(i))
		time.Sleep(time.Millisecond * 100)
	}
	bar.Finish()
}
