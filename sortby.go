package howlongtobeat

type SortBy string

const (
	SortByName            SortBy = "name"
	SortByMainStory       SortBy = "main"
	SortByMainPlusExtra   SortBy = "mainp"
	SortByCompletionist   SortBy = "comp"
	SortByAverageTime     SortBy = "averagea"
	SortByTopRated        SortBy = "rating"
	SortByMostPopular     SortBy = "popular"
	SortByMostBacklogs    SortBy = "backlog"
	SortByMostSubmissions SortBy = "usersp"
	SortByMostPlayed      SortBy = "playing"
	SortByMostSpeedruns   SortBy = "speedruns"
	SortByMostReviews     SortBy = "reviews"
	SortByReleaseDate     SortBy = "release"
)
