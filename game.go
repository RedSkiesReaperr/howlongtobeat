package howlongtobeat

type Game struct {
	Id                   int    `json:"game_id"`
	Name                 string `json:"game_name"`
	NameDate             int    `json:"game_name_date"`
	Alias                string `json:"game_alias"`
	Type                 string `json:"game_type"`
	Image                string `json:"game_image"`
	CompletionLvlCombine int    `json:"comp_lvl_combine"`
	CompletionLvlSp      int    `json:"comp_lvl_sp"`
	CompletionLvlCo      int    `json:"comp_lvl_co"`
	CompletionLvlMp      int    `json:"comp_lvl_mp"`
	CompletionLvlSpd     int    `json:"comp_lvl_spd"`
	CompletionMain       int    `json:"comp_main"`
	CompletionPlus       int    `json:"comp_plus"`
	CompletionFull       int    `json:"comp_100"`
	CompletionAll        int    `json:"comp_all"`
	CompletionMainCount  int    `json:"comp_main_count"`
	CompletionPlusCount  int    `json:"comp_plus_count"`
	CompletionFullCount  int    `json:"comp_100_count"`
	CompletionAllCount   int    `json:"comp_all_count"`
	InvestedCo           int    `json:"invested_co"`
	InvestedMp           int    `json:"invested_mp"`
	InvestedCoCount      int    `json:"invested_co_count"`
	InvestedMpCount      int    `json:"invested_mp_count"`
	CountCompletion      int    `json:"count_comp"`
	CountSpeedrun        int    `json:"count_speedrun"`
	CountBacklog         int    `json:"count_backlog"`
	CountReview          int    `json:"count_review"`
	ReviewScore          int    `json:"review_score"`
	CountPlaying         int    `json:"count_playing"`
	CountRetired         int    `json:"count_retired"`
	ProfileDev           string `json:"profile_dev"`
	ProfilePopular       int    `json:"profile_popular"`
	ProfileSteam         int    `json:"profile_steam"`
	ProfilePlatform      string `json:"profile_platform"` //TODO: Parse by ',' + type []string
	ReleaseWord          int    `json:"release_world"`
}
