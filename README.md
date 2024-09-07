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
