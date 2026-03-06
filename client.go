package howlongtobeat

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/corpix/uarand"
)

const searchIdTimeout int = 2 // Hours

type Client struct {
	api             *api
	apiInfosFoundAt time.Time
	authToken       string
}

func New() (*Client, error) {
	client := &Client{}

	if err := client.findApiInfos(); err != nil {
		return nil, fmt.Errorf("can't find api infos: %v", err)
	}

	if err := client.refreshAuthToken(); err != nil {
		fmt.Printf("WARNING: can't get auth token: %v\n", err)
	}

	return client, nil
}

func (c *Client) findApiInfos() error {
	api, err := scrapApiInfos()
	if err != nil {
		return fmt.Errorf("scrap: %s", err)
	}

	c.api = &api
	c.apiInfosFoundAt = time.Now()

	return nil
}

func (c Client) searchApiInfosTimedOut() bool {
	return time.Since(c.apiInfosFoundAt).Hours() >= float64(searchIdTimeout)
}

func (c *Client) refreshAuthToken() error {
	// Fallback path if missing
	apiPath := "/api/finder"
	if c.api != nil && c.api.path != "" {
		apiPath = c.api.path
	}

	initUrl := fmt.Sprintf("https://howlongtobeat.com%s/init?t=%d", apiPath, time.Now().UnixMilli())

	req, _ := http.NewRequest("GET", initUrl, nil)
	req.Header = map[string][]string{
		"Accept":     {"*/*"},
		"User-Agent": {uarand.GetRandom()},
		"Referer":    {"https://howlongtobeat.com/"},
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	body, _ := io.ReadAll(resp.Body)

	var result struct {
		Token string `json:"token"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return fmt.Errorf("unmarshal: %v (Body: %s)", err, string(body))
	}

	c.authToken = result.Token
	return nil
}
