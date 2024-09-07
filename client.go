package howlongtobeat

import "fmt"

type Client struct {
	searchId string
}

func New() (*Client, error) {
	client := &Client{}
	foundId, err := client.findSearchId()
	if err != nil {
		return nil, fmt.Errorf("can't find searchId: %v", err)
	}

	client.searchId = foundId

	return client, nil
}

func (c *Client) findSearchId() (string, error) {
	//TODO: Search dynamic ID
	return "3ef777b4d4ae0be", nil
}
