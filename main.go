package main

import (
	"github.com/labstack/gommon/log"
	"github.com/sclevine/agouti"
)

func main() {
	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		log.Fatalf("driver start failed %#v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		log.Fatalf("new page failed %#v", err)
	}

	if err := page.Navigate("http://qiita.com/"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}
	err = page.Screenshot("./chrome_qiita.jpg")
	if err != nil {
		log.Fatalf("page screenshot failed %#v", err)
	}

	page.FindByLink("もっと詳しく").Click()
	log.Info(page.Title())
	err = page.Screenshot("./chrome_qiita2.jpg")
	if err != nil {
		log.Fatalf("page screenshot2 failed %#v", err)
	}
}
