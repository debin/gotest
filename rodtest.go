package main

import (
	"fmt"
	_ "os"
	"time"

	"github.com/go-rod/rod"
)

func main() {
	// Launch a new browser with default options, and connect to it.
	browser := rod.New().MustConnect()

	// Even you forget to close, rod will close it after main process ends.
	defer browser.MustClose()

	// Timeout will be passed to all chained function calls.
	// The code will panic out if any chained call is used after the timeout.
	page := browser.Timeout(time.Minute).MustPage("https://www.zhidaohu.com")

	page.WaitRequestIdle(3*time.Millisecond,nil,nil)

	//wait()

	text := page.MustElement("html").MustText()

	fmt.Println("result is:")
	fmt.Println(text)


	fmt.Println("run over")

	// Output:
	// Git is the most widely used version control system.
	// Found 5 input elements
	// 1 + 2 = 3
	// Search · git · GitHub
}
