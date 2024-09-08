package howlongtobeat

type Modifier string

const (
	ModifierAll         Modifier = ""
	ModifierHideDlc     Modifier = "hide_dlc"
	ModifierOnlyDlc     Modifier = "only_dlc"
	ModifierOnlyMods    Modifier = "only_mods"
	ModifierOnlyHacks   Modifier = "only_hacks"
	ModifierHiddenStats Modifier = "hidden_stats"
	ModifierUserStats   Modifier = "user_stats"
)
