package howlongtobeat

import (
	"fmt"
	"time"
)

const searchIdTimeout int = 6 // Hours

type Client struct {
	searchId        string
	searchIdFoundAt time.Time
}

func New() (*Client, error) {
	client := &Client{}

	if err := client.findSearchId(); err != nil {
		return nil, fmt.Errorf("can't find searchId: %v", err)
	}

	return client, nil
}

func (c *Client) findSearchId() error {
	id, err := scrapSearchId()
	if err != nil {
		return fmt.Errorf("scrap: %s", err)
	}

	c.searchId = id
	c.searchIdFoundAt = time.Now()

	return nil
}

func (c Client) searchIdTimedOut() bool {
	return time.Since(c.searchIdFoundAt).Hours() >= float64(searchIdTimeout)
}
