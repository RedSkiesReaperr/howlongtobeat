package howlongtobeat

import (
	"fmt"
	"time"
)

const searchIdTimeout int = 2 // Hours

type Client struct {
	api             *api
	apiInfosFoundAt time.Time
}

func New() (*Client, error) {
	client := &Client{}

	if err := client.findApiInfos(); err != nil {
		return nil, fmt.Errorf("can't find api infos: %v", err)
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
