package howlongtobeat

import "fmt"

type api struct {
	path string
	err  error
}

func (a api) searchUrl() string {
	base := "https://howlongtobeat.com"

	return fmt.Sprintf("%s%s", base, a.path)
}
