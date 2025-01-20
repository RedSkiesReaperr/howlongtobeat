package howlongtobeat

type Perspective string

const (
	PerspectiveAll            Perspective = ""
	PerspectiveFirstPerson    Perspective = "First-Person"
	PerspectiveIsometric      Perspective = "Isometric"
	PerspectiveSide           Perspective = "Side"
	PerspectiveText           Perspective = "Text"
	PerspectiveThirdPerson    Perspective = "Third-Person"
	PerspectiveTopDown        Perspective = "Top-Down"
	PerspectiveVirtualReality Perspective = "Virtual Reality"
)

type Flow string

const (
	FlowAll                  Flow = ""
	FlowIncremental          Flow = "Incremental"
	FlowMassivelyMultiplayer Flow = "Massively Multiplayer"
	FlowMultidirectional     Flow = "Multidirectional"
	FlowOnRails              Flow = "On-Rails"
	FlowPointAndClick        Flow = "Point-and-Click"
	FlowRealTime             Flow = "Real-Time"
	FlowScrolling            Flow = "Scrolling"
	FlowTurnBased            Flow = "Turn-Based"
)

type Genre string

const (
	GenreAll              Genre = ""
	GenreAction           Genre = "Action"
	GenreAdventure        Genre = "Adventure"
	GenreArcade           Genre = "Arcade"
	GenreBattleArena      Genre = "Battle Arena"
	GenreBeatEmUp         Genre = "Beat em Up"
	GenreBoardGame        Genre = "Board Game"
	GenreBreakout         Genre = "Breakout"
	GenreCardGame         Genre = "Card Game"
	GenreCityBuilding     Genre = "City-Building"
	GenreCompilation      Genre = "Compilation"
	GenreEducational      Genre = "Educational"
	GenreFighting         Genre = "Fighting"
	GenreFitness          Genre = "Fitness"
	GenreFlight           Genre = "Flight"
	GenreFullMotionVideo  Genre = "Full Motion Video (FMV)"
	GenreHackAndSlash     Genre = "Hack and Slash"
	GenreHiddenObject     Genre = "Hidden Object"
	GenreHorror           Genre = "Horror"
	GenreInteractiveArt   Genre = "Interactive Art"
	GenreManagement       Genre = "Management"
	GenreMusicAndRhythm   Genre = "Music/Rhythm"
	GenreOpenWorld        Genre = "Open World"
	GenreParty            Genre = "Party"
	GenrePinball          Genre = "Pinball"
	GenrePlatform         Genre = "Platform"
	GenrePuzzle           Genre = "Puzzle"
	GenreRacingAndDriving Genre = "Racing/Driving"
	GenreRoguelike        Genre = "Roguelike"
	GenreRolePlaying      Genre = "Role-Playing"
	GenreSandbox          Genre = "Sandbox"
	GenreShooter          Genre = "Shooter"
	GenreSimulation       Genre = "Simulation"
	GenreSocial           Genre = "Social"
	GenreSports           Genre = "Sports"
	GenreStealth          Genre = "Stealth"
	GenreStrategyTactical Genre = "Strategy/Tactical"
	GenreSurvival         Genre = "Survival"
	GenreTowerDefense     Genre = "Tower Defense"
	GenreTrivia           Genre = "Trivia"
	GenreVehicularCombat  Genre = "Vehicular Combat"
	GenreVisualNovel      Genre = "Visual Novel"
)

type Difficulty string

const (
	DifficultyAll Difficulty = ""
)
