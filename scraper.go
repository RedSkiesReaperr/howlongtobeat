package howlongtobeat

import (
	"strings"
	"time"

	"github.com/go-rod/rod"
)

const url = "https://howlongtobeat.com/?q=elden%2520ring"

func scrapSearchId() (string, error) {
	foundSearchId := ""
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	router := browser.HijackRequests()
	defer router.MustStop()

	router.MustAdd("*/api/search/*", func(ctx *rod.Hijack) {
		result := strings.Split(ctx.Request.URL().String(), "/")
		foundSearchId = result[len(result)-1]
	})
	go router.Run()

	browser.MustPage(url).MustWaitElementsMoreThan("input[name=site-search]", 0)

	waitSeconds := 10
	for i := 0; i < waitSeconds && foundSearchId == ""; i++ {
		time.Sleep(1 * time.Second)
	}

	return foundSearchId, nil
}
