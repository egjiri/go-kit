package progress

import (
	"errors"
	"fmt"
	"time"

	"github.com/gosuri/uiprogress"
)

// barManager manages the Progress of the bars
var barManager = uiprogress.New()

// Bar managest the state of a progres bar
type bar struct {
	bar *uiprogress.Bar
}

// NewBar creates a prgoress bar with reasonable defaults, configurable count
// and action to display in the suffix of AppendFunc
func NewBar(count int, action string) (*bar, error) {
	if count < 1 {
		return nil, errors.New("invalid progress bar count")
	}

	b := barManager.AddBar(count)
	progressBar := bar{b}
	startTime := time.Now()
	var duration string
	b.PrependFunc(func(b *uiprogress.Bar) string {
		if progressBar.InProgress() {
			duration = time.Since(startTime).Round(time.Second).String()
		}
		return duration
	})
	b.AppendFunc(func(b *uiprogress.Bar) string {
		return fmt.Sprintf("%.1f%% - %v %v/%v", b.CompletedPercent(), action, b.Current(), b.Total)
	})
	return &progressBar, nil
}

// Incr increments the progress bar
func (b *bar) Incr() bool {
	return b.bar.Incr()
}

// InProgress returns whether the progress bar still has more to increment through
func (b *bar) InProgress() bool {
	return b.bar.CompletedPercent() < 100
}

// String returns the string representation of the progress bar
func (b *bar) String() string {
	return b.bar.String()
}

// Start begins the rendering of the progress bars
func Start() {
	barManager.Start()
}
