package progress

import (
	"fmt"
	"time"

	"github.com/gosuri/uiprogress"
)

// BarManager manages the Progress of the bars
var BarManager = uiprogress.New()

// Bar creates a prgoress bar with reasonable defaults, configurable count
// and action to display in the suffix of AppendFunc
func Bar(count int, action string) (*uiprogress.Bar, error) {
	if count < 1 {
		return nil, fmt.Errorf("invalid progress bar count %v", count)
	}

	bar := BarManager.AddBar(count)
	startTime := time.Now()
	var duration string
	bar.PrependFunc(func(b *uiprogress.Bar) string {
		if b.CompletedPercent() < 100 {
			duration = time.Since(startTime).Round(time.Second).String()
		}
		return duration
	})
	bar.AppendFunc(func(b *uiprogress.Bar) string {
		return fmt.Sprintf("%.1f%% - %v %v/%v", b.CompletedPercent(), action, b.Current(), b.Total)
	})
	return bar, nil
}
