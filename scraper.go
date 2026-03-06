package howlongtobeat

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
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
	baseUrl := "https://howlongtobeat.com/"
	result := scrapResult{}

	browser := setupBrowser()
	defer browser.MustClose()

	router, err := setupRouter(browser)
	if err != nil {
		return api{}, fmt.Errorf("can't setup router: %v", err)
	}
	defer router.MustStop()

	if err := router.Add("*.js", proto.NetworkResourceTypeScript, result.process); err != nil {
		return api{}, fmt.Errorf("hijack failed: %v", err)
	}

	go router.Run()
	page := browser.MustPage(baseUrl)

	// Scroll to trigger JS
	page.MustEval("() => window.scrollTo(0, document.body.scrollHeight)")
	page.MustWaitElementsMoreThan("input[name=site-search]", 0)

	waitSeconds := 15
	for i := 0; i < waitSeconds && !result.isScrapped(); i++ {
		time.Sleep(1 * time.Second)
	}

	return api{path: result.apiPath}, result.scrapErr
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
	err := ctx.LoadResponse(http.DefaultClient, true)
	if err != nil {
		return
	}
	scriptContent := ctx.Response.Body() // Script content

	path, err := extractApiInfos(scriptContent)
	if err != nil {
		return
	}

	sr.apiPath = path
}

func (sr *scrapResult) isScrapped() bool {
	return sr.apiPath != "" || sr.scrapErr != nil
}

func extractApiInfos(src string) (path string, err error) {
	// Pattern to find the search endpoint.
	re := regexp.MustCompile(`(?i)fetch\s*\(\s*["'](\/api\/[a-zA-Z0-9_/]+)[^"']*["']\s*,\s*\{[^}]*method:\s*["']POST["']`)
	match := re.FindStringSubmatch(src)

	if match != nil {
		apiPath := match[1]

		// Ignore non-search endpoints
		if strings.Contains(apiPath, "logout") || strings.Contains(apiPath, "user") || strings.Contains(apiPath, "error") {
			return "", fmt.Errorf("ignoring non-search endpoint: %s", apiPath)
		}

		return apiPath, nil
	}

	return "", fmt.Errorf("no valid match found")
}
