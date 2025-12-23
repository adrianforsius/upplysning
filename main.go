package main

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/go-rod/rod/lib/utils"
	"github.com/ysmood/gson"
)

func main() {
	// url := "https://www.upplysning.se/person/?x=0791&f=christina&c=vaxholm&ac=0187&county=01&municipality=0187&malegender=False&femalegender=True&m=0&sl=detail&page=2"
	url := "https://www.wikipedia.org/"
	launcher := launcher.New().Bin("google-chrome-stable").Headless(false).MustLaunch()

	browser := rod.New().ControlURL(launcher).MustConnect()
	defer browser.Close()

	// The 2 lines below share the same context, they will be canceled after 2 seconds in total
	// capture entire browser viewport, returning jpg with quality=90
	img, err := browser.MustPage(url).MustWaitStable().ScrollScreenshot(&rod.ScrollScreenshotOptions{
		Format:  proto.PageCaptureScreenshotFormatJpeg,
		Quality: gson.Int(90),
	})
	if err != nil {
		panic(err)
	}

	_ = utils.OutputFile("my.jpg", img)
}
