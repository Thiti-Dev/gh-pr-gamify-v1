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

	appConf, err := config.LoadApplicationConfig()
	if err != nil {
		log.Fatalf("Failed to load application config: %v", err)
	}

	fetchers := []fetcher.FetcherI{}

	for _, repo := range appConf.Github.Repositories {
		fetchers = append(fetchers, fetcher.NewFetcher(requester.NewRestyRequester(), fetcher.RepositoryPointer{Token: cfg.GithubBearerToken, Organiztion: repo.Organization, Repository: repo.Name}))
	}

	runner := runner.NewRunner(fetchers, cfg)
	err = runner.Run()
	if err != nil {
		log.Fatalf("Failed to run: %v", err)
	}
}
