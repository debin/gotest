package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true),
		chromedp.Flag("hide-scrollbars", false),
		chromedp.Flag("mute-audio", false),
		chromedp.ExecPath("C:/Program Files (x86)/Microsoft/Edge/Application/msedge.exe"),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}

	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	c, cc := chromedp.NewExecAllocator(ctx, options...)
	defer cc()
	// create context
	ctx, cancel := chromedp.NewContext(c)
	defer cancel()

	var content string

	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.baidu.com`),
		chromedp.Text(`body`, &content, chromedp.NodeVisible, chromedp.ByQuery),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("content:")
	log.Println(content)


}