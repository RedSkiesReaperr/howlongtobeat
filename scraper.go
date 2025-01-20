package howlongtobeat

import (
	"fmt"
	"regexp"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

type scrapResult struct {
	apiPath  string
	apiKey   string
	scrapErr error
}

func scrapApiInfos() (api, error) {
	baseUrl := "https://howlongtobeat.com/?q=elden%2520ring"
	result := scrapResult{}

	browser := setupBrowser()
	defer browser.MustClose()

	router, err := setupRouter(browser)
	if err != nil {
		return api{}, fmt.Errorf("can't setup router: %v", err)
	}
	defer router.MustStop()

	if err := router.Add("*_app-*.js", proto.NetworkResourceTypeScript, result.process); err != nil {
		return api{}, fmt.Errorf("hijack failed: %v", err)
	}

	go router.Run()
	browser.MustPage(baseUrl).MustWaitElementsMoreThan("input[name=site-search]", 0)

	waitSeconds := 10
	for i := 0; i < waitSeconds && !result.isScrapped(); i++ {
		time.Sleep(1 * time.Second)
	}

	return api{key: result.apiKey, path: result.apiPath}, result.scrapErr
}

func setupBrowser() *rod.Browser {
	path, _ := launcher.LookPath()
	u := launcher.New().Bin(path).MustLaunch()
	browser := rod.New().ControlURL(u).MustConnect()

	return browser
}

func setupRouter(browser *rod.Browser) (*rod.HijackRouter, error) {
	router := browser.HijackRequests()

	return router, nil
}

func (sr *scrapResult) process(ctx *rod.Hijack) {
	ctx.MustLoadResponse()               // Force to continue script downloading reaquest
	scriptContent := ctx.Response.Body() // Script content

	path, key, err := extractApiInfos(scriptContent)
	if err != nil {
		sr.scrapErr = fmt.Errorf("extract api infos: %v", err)
		return
	}

	sr.apiKey = key
	sr.apiPath = path
}

func (sr *scrapResult) isScrapped() bool {
	return sr.apiPath != "" || sr.apiKey != "" || sr.scrapErr != nil
}

func extractApiInfos(src string) (path string, key string, err error) {
	pattern := `(\/api\/[a-zA-Z0-9\/]*)".concat\("([a-zA-Z0-9]+)"\).concat\("([a-zA-Z0-9]+)"\)`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(src)

	if matches == nil {
		return "", "", fmt.Errorf("no match found")
	}

	path = matches[1]
	key = matches[2] + matches[3]
	err = nil
	return
}
