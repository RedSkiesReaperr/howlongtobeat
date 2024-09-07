package howlongtobeat

import "fmt"

type SearchType string

const (
	SearchTypeGames SearchType = "games"
)

type SearchRequest struct {
	SearchType    string        `json:"searchType,required"`
	SearchTerms   []string      `json:"searchTerms,required"`
	SearchPage    int           `json:"searchPage,required"`
	PageSize      int           `json:"size,required"`
	SearchOptions searchOptions `json:"searchOptions,required"`
}

type searchOptions struct {
	Games      searchOptionsGames `json:"games,required"`
	Users      searchOptionsUsers `json:"users,required"`
	Filter     string             `json:"filter,required"`
	Sort       int                `json:"sort,required"`
	Randomizer int                `json:"randomizer,required"`
}

type searchOptionsGames struct {
	UserId        int                         `json:"userId,required"`
	Platform      string                      `json:"platform,required"`
	SortCategory  string                      `json:"sortCategory,required"`
	RangeCategory string                      `json:"rangeCategory,required"`
	RangeTime     searchOptionsGamesRangeTime `json:"rangeTime,required"`
	Gameplay      searchOptionsGamesGameplay  `json:"gameplay,required"`
	Modifier      string                      `json:"modifier,required"`
}

type searchOptionsUsers struct {
	SortCategory string `json:"sortCategory"`
}

type searchOptionsGamesRangeTime struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type searchOptionsGamesGameplay struct {
	Perspective string `json:"perspective"`
	Flow        string `json:"flow"`
	Genre       string `json:"genre"`
}

type SearchResult struct {
	Color       string `json:"color"`
	Title       string `json:"title"`
	Category    string `json:"category"`
	Count       int    `json:"count"`
	PageCurrent int    `json:"pageCurrent"`
	PageTotal   int    `json:"pageTotal"`
	PageSize    int    `json:"pageSize"`
	Data        []Game `json:"data"`
}

func (c Client) Search(request SearchRequest) (SearchResult, error) {
	result, err := request.send(c.searchId)
	if err != nil {
		return SearchResult{}, fmt.Errorf("request failed: %v", err)
	}

	return result, nil
}
