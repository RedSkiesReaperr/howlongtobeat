# howlongtobeat
Unofficial Go implementation of HowLongToBeat website API

# Installation
```bash
$ go get github.com/RedSkiesReaperr/howlongtobeat
```

# Example
```go
hltb, err := howlongtobeat.New()
if err != nil {
    // error management
}

request, err := howlongtobeat.NewSearchRequest("elden ring")
if err != nil {
    // error management
}

result, err := hltb.Search(request)
if err != nil {
    // error management
}

for _, game := range result.Data {
    fmt.Println(game.Name)
}
```
# Customize your request
After creating a SearchRequest you can send it as is or customize it, to do so you have several options: 

#### Customize search terms
```go
request.SetSearchTerms("new terms")
```

#### Customize target platform
You can find a list of available platforms [here](https://github.com/RedSkiesReaperr/howlongtobeat/blob/main/platforms.go)
```go
request.SetPlatform(howlongtobeat.PlatformPC)
```

#### Customize pagination
```go
// Retrieve page number 2 with a page size of 25
request.SetPagination(2, 25)
```

#### Customize game modifiers
You can find a list of available modifiers [here](https://github.com/RedSkiesReaperr/howlongtobeat/blob/main/modifiers.go)
```go
request.SetModifier(howlongtobeat.ModifierOnlyDlc)
```

#### Customize sorting
You can find a list of available sort kinds [here](https://github.com/RedSkiesReaperr/howlongtobeat/blob/main/sortby.go)
```go
request.SetSorting(howlongtobeat.SortByReleaseDate)
```

#### Customize gameplay
You can find a list of available gameplay parts [here](https://github.com/RedSkiesReaperr/howlongtobeat/blob/main/gameplay.go)
```go
request.SetGameplay(howlongtobeat.PerspectiveThirdPerson, howlongtobeat.FlowPointAndClick, howlongtobeat.GenreHorror, howlongtobeat.DifficultyAll)
```

## Dependencies
Thanks to all the authors who created and maintains the following packages:
- [go-rod/rod](https://github.com/go-rod/rod)
- [corpix/uarand](https://github.com/corpix/uarand)

## Contributing
If you wish to contribute to this project, please follow these steps:
1. Fork this repository.
2. Create a branch for your feature: git checkout -b feature/feature-name
3. Commit your changes: git commit -m 'Added a new feature'
4. Push your branch: git push origin feature/feature-name
5. Open a Pull Request.

## License
This project is licensed under the MIT License. See the LICENSE file for details.
Feel free to open issues or submit feature requests if you have ideas to enhance this project.
