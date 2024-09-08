package howlongtobeat

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/corpix/uarand"
)

type SearchRequest struct {
	SearchType    SearchType    `json:"searchType,required"`
	SearchTerms   []string      `json:"searchTerms,required"`
	SearchPage    int           `json:"searchPage,required"`
	PageSize      int           `json:"size,required"`
	SearchOptions searchOptions `json:"searchOptions,required"`
}

func NewSearchRequest(searchTerms string) (SearchRequest, error) {
	request := SearchRequest{}
	request.setDefaults()
	request.SetSearchTerms(searchTerms)

	return request, nil
}

func (s *SearchRequest) SetSearchTerms(terms string) {
	result := strings.Split(terms, " ")

	s.SearchTerms = result
}

func (s *SearchRequest) SetPlatform(newP Platform) {
	s.SearchOptions.Games.Platform = newP
}

func (s *SearchRequest) SetPagination(pageIndex, pageSize int) {
	s.SearchPage = pageIndex
	s.PageSize = pageSize
}

func (s *SearchRequest) SetModifier(modifier Modifier) {
	s.SearchOptions.Games.Modifier = modifier
}

func (s *SearchRequest) SetSorting(sort SortBy) {
	s.SearchOptions.Games.SortCategory = sort
}

func (s *SearchRequest) SetGameplay(pers Perspective, flow Flow, genre Genre) {
	s.SearchOptions.Games.Gameplay.Perspective = pers
	s.SearchOptions.Games.Gameplay.Flow = flow
	s.SearchOptions.Games.Gameplay.Genre = genre
}

func (s *SearchRequest) setDefaults() {
	s.SearchType = SearchTypeGames
	s.SearchPage = 1
	s.PageSize = 20
	s.SearchOptions.Games = searchOptionsGames{
		UserId:        0,
		Platform:      PlatformAll,
		SortCategory:  SortByMostPopular,
		RangeCategory: "main",
		RangeTime:     searchOptionsGamesRangeTime{},
		Gameplay: searchOptionsGamesGameplay{
			Perspective: PerspectiveAll,
			Flow:        FlowAll,
			Genre:       GenreAll,
		},
		Modifier: ModifierAll,
	}
	s.SearchOptions.Users = searchOptionsUsers{
		SortCategory: "postcount",
	}
	s.SearchOptions.Filter = ""
	s.SearchOptions.Sort = 0
	s.SearchOptions.Randomizer = 0
}

func (s SearchRequest) send(searchId string) (SearchResult, error) {
	url := fmt.Sprintf("https://howlongtobeat.com/api/search/%s", searchId)
	result := SearchResult{}
	httpClient := &http.Client{}

	requestBody, err := json.Marshal(s)
	if err != nil {
		return result, fmt.Errorf("invalid body: %s", err)
	}

	req, _ := http.NewRequest("POST", url, strings.NewReader(string(requestBody)))
	req.Header = map[string][]string{
		"Content-Type": {"application/json"},
		"User-Agent":   {uarand.GetRandom()},
		"Referer":      {"https://howlongtobeat.com/"},
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return result, fmt.Errorf("error: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, fmt.Errorf("can't read body: %s", err)
	}

	if resp.StatusCode != 200 {
		return result, fmt.Errorf("request failed: %s", resp.Status)
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, fmt.Errorf("invalid response body: %s", err)
	}

	return result, nil
}
