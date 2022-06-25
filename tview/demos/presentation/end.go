package main

import (
	"fmt"

	tcell "github.com/gdamore/tcell/v2"
	"github.com/5qw/XyzCord/tview"
)

// End shows the final slide.
func End(nextSlide func()) (title string, content tview.Primitive) {
	textView := tview.NewTextView().SetDoneFunc(func(key tcell.Key) {
		nextSlide()
	})
	url := "https://github.com/5qw/XyzCord/tview"
	fmt.Fprint(textView, url)
	return "End", Center(len(url), 1, textView)
}
