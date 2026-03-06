package main

import (
	"fmt"
	"log"

	"github.com/RedSkiesReaperr/howlongtobeat"
)

func main() {
	// 1. Create a new Client
	hltb, err := howlongtobeat.New()
	if err != nil {
		log.Fatalf("can't find api infos: %v", err)
	}

	// 2. Create a search request for "Elden Ring" (as an example)
	gameName := "Elden Ring"
	request, err := howlongtobeat.NewSearchRequest(gameName)
	if err != nil {
		log.Fatalf("can't create search request: %v", err)
	}

	// Optional: customize your request
	// request.SetPlatform(howlongtobeat.PlatformPC)
	// request.SetSorting(howlongtobeat.SortByMostPopular)

	// 3. Perform the search
	fmt.Printf("Searching for: %s...\n", gameName)
	result, err := hltb.Search(request)
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}

	// 4. Display results
	fmt.Printf("Found %d results (Page %d/%d):\n\n", result.Count, result.PageCurrent, result.PageTotal)

	for _, game := range result.Data {
		fmt.Printf("--- %s (ID: %d) ---\n", game.Name, game.Id)
		fmt.Printf("Main Story: %.2f hours\n", float64(game.CompletionMain)/3600.0)
		fmt.Printf("Main + Extra: %.2f hours\n", float64(game.CompletionPlus)/3600.0)
		fmt.Printf("Completionist: %.2f hours\n", float64(game.CompletionFull)/3600.0)
		fmt.Printf("Review Score: %d/100\n", game.ReviewScore)
		fmt.Println()
	}
}
