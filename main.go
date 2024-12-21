package main

import (
	"log"

	"github.com/Thiti-Dev/gh-pr-gamify-v1/core/fetcher"
	"github.com/Thiti-Dev/gh-pr-gamify-v1/core/runner"
	"github.com/Thiti-Dev/gh-pr-gamify-v1/pkg/config"
	"github.com/Thiti-Dev/gh-pr-gamify-v1/pkg/requester"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// TODO: load from the config
	smarterMailFetcher := fetcher.NewFetcher(requester.NewRestyRequester(), fetcher.RepositoryPointer{Token: cfg.GithubBearerToken, Organiztion: "smartertravel", Repository: "partner-feed"})

	runner := runner.NewRunner([]fetcher.FetcherI{smarterMailFetcher})
	err = runner.Run()
	if err != nil {
		log.Fatalf("Failed to run: %v", err)
	}
}
