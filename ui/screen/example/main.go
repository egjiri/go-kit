package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/egjiri/go-kit/ui/progress"
	"github.com/egjiri/go-kit/ui/screen"
	"github.com/fatih/color"
)

type tracker struct{}

func (t *tracker) String() string {
	return fmt.Sprintf(`JOB IN PROGRESS
========================
Valid Records Saved:   %v
Invalid Records Saved: %v
Duplicates Skipped:    %v
Entries Updated:       %v
Time Elapsed:          %v
`, r(), r(), r(), r(), r())
}

func r() string {
	i := rand.Int() % 15
	s := strconv.Itoa(i)
	return color.CyanString(s)
}

func main() {
	// screen.Add("ENDRI")
	// screen.Add("Gjiri")
	// screen.Refresh()

	// bar, _ := progress.NewBar(10, "Geocoded")
	// bar2, _ := progress.NewBar(20, "Imported")
	// for bar2.InProgress() {
	// 	time.Sleep(time.Second)
	// 	bar.Incr()
	// 	bar2.Incr()
	// }

	// for {
	// 	time.Sleep(time.Second)
	// 	screen.Println("Warning: Geocoding Failed!")
	// }

	screen.Add("Endri Gjiri")
	screen.Add("Gjiri\nEndri\n")
	bar, _ := progress.NewBar(10, "Geocoded")
	bar2, _ := progress.NewBar(20, "Imported")
	screen.Add("", new(tracker))
	for bar2.InProgress() {
		time.Sleep(time.Second)
		screen.Println("Warning: Geocoding Failed!")
		bar.Incr()
		bar2.Incr()
	}
}
