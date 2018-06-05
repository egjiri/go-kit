package progress

import (
	"fmt"
	"time"

	"github.com/egjiri/go-kit/ui/screen"
	"github.com/gosuri/uiprogress"
	"github.com/pkg/errors"
)

// barManager manages the Progress of the bars
var barManager = uiprogress.New()

// Bar managest the state of a progres bar
type Bar struct {
	bar *uiprogress.Bar
}

// NewBar creates a prgoress bar with reasonable defaults, configurable count
// and action to display in the suffix of AppendFunc
func NewBar(count int, action string) (*Bar, error) {
	if count < 1 {
		return nil, errors.New("invalid progress bar count")
	}
	bar := Bar{barManager.AddBar(count)}
	addBarDecorators(&bar, count, action)
	screen.Add(&bar) // Add progress bar to the screen
	return &bar, nil
}

// Incr increments the progress bar
func (b *Bar) Incr() bool {
	res := b.bar.Incr()
	screen.Refresh() // Refresh screen every time the progress bar increments
	return res
}

// InProgress returns whether the progress bar still has more to increment through
func (b *Bar) InProgress() bool {
	return b.bar.CompletedPercent() < 100
}

// String returns the string representation of the progress bar
func (b *Bar) String() string {
	return b.bar.String()
}

func addBarDecorators(bar *Bar, count int, action string) {
	startTime := time.Now()
	var duration string
	bar.bar.PrependFunc(func(b *uiprogress.Bar) string {
		if bar.InProgress() {
			duration = time.Since(startTime).Round(time.Second).String()
		}
		return duration
	})
	bar.bar.AppendFunc(func(b *uiprogress.Bar) string {
		str := fmt.Sprintf("%.1f%%", b.CompletedPercent())
		if b.CompletedPercent() == 0 {
			str += "%"
		}
		return str + fmt.Sprintf(" - %v %v/%v", action, b.Current(), b.Total)
	})
}
