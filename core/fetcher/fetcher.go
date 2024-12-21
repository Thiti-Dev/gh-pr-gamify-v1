package fetcher

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Thiti-Dev/gh-pr-gamify-v1/models"
	"github.com/Thiti-Dev/gh-pr-gamify-v1/pkg/requester"
)

type FetcherI interface {
	GetPullRequestList(dRange models.DateRange) (*models.PRResponse, error)
	GetMergedPullRequestList(dRange models.DateRange) (*models.PRResponse, error)
	GetFormattedRepositoryPath() string
	GetPullRequestReviews(id int) ([]models.PRReview, error)
}

type RepositoryPointer struct {
	Organiztion string
	Repository  string
	Token       string
}

type Fetcher struct {
	client            requester.RequesterI
	repositoryPointer RepositoryPointer
}

func NewFetcher(client requester.RequesterI, rp RepositoryPointer) FetcherI {
	return &Fetcher{
		client:            client,
		repositoryPointer: rp,
	}
}
func (f *Fetcher) GetPullRequestReviews(id int) ([]models.PRReview, error) {
	req, err := f.client.Get(context.Background(), fmt.Sprintf("https://api.github.com/repos/%s/%s/pulls/%d/reviews", f.repositoryPointer.Organiztion, f.repositoryPointer.Repository, id), map[string]string{
		"Authorization": "Bearer " + f.repositoryPointer.Token,
		"Accept":        "application/vnd.github.v3+json",
	})
	if err != nil {
		return nil, err
	}

	var res []models.PRReview
	err = json.Unmarshal(req.Body, &res)

	return res, err
}

func (f *Fetcher) GetPullRequestList(dRange models.DateRange) (*models.PRResponse, error) {
	rangeParams := fmt.Sprintf("%s..%s", dRange.GetFormattedFrom(), dRange.GetFormattedTo())
	req, err := f.client.Get(context.Background(), fmt.Sprintf("https://api.github.com/search/issues?q=repo:%s/%s+is:pr+created:%s", f.repositoryPointer.Organiztion, f.repositoryPointer.Repository, rangeParams), map[string]string{
		"Authorization": "Bearer " + f.repositoryPointer.Token,
		"Accept":        "application/vnd.github.v3+json",
	})
	if err != nil {
		return nil, err
	}

	var res models.PRResponse
	err = json.Unmarshal(req.Body, &res)

	return &res, err
}

func (f *Fetcher) GetMergedPullRequestList(dRange models.DateRange) (*models.PRResponse, error) {
	rangeParams := fmt.Sprintf("%s..%s", dRange.GetFormattedFrom(), dRange.GetFormattedTo())
	req, err := f.client.Get(context.Background(), fmt.Sprintf("https://api.github.com/search/issues?q=repo:%s/%s+is:pr+merged:%s", f.repositoryPointer.Organiztion, f.repositoryPointer.Repository, rangeParams), map[string]string{
		"Authorization": "Bearer " + f.repositoryPointer.Token,
		"Accept":        "application/vnd.github.v3+json",
	})
	if err != nil {
		return nil, err
	}

	var res models.PRResponse
	err = json.Unmarshal(req.Body, &res)

	return &res, err
}

func (f *Fetcher) GetFormattedRepositoryPath() string {
	return fmt.Sprintf("%s/%s", f.repositoryPointer.Organiztion, f.repositoryPointer.Repository)
}
