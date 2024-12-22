package services

import (
	"fmt"
	"time"

	"github.com/Thiti-Dev/gh-pr-gamify-v1/core/fetcher"
	"github.com/Thiti-Dev/gh-pr-gamify-v1/models"
	prState "github.com/Thiti-Dev/gh-pr-gamify-v1/types/pr-state"
)

type PRService struct {
	OperatedTime time.Time
	Ent          *models.PRItem
}

type PRServiceOption func(*PRService)

func WithPRItem(ent models.PRItem) PRServiceOption {
	return func(ps *PRService) {
		ps.Ent = &ent
	}
}

// Functional Options Pattern
/*
	prServiceWithEnt := NewPRService(time.Now(), WithPRItem(models.PRItem{}))
	prServiceWithoutEnt := NewPRService(time.Now())
*/
func NewPRService(operatedTime time.Time, opts ...PRServiceOption) *PRService {
	pr := &PRService{
		OperatedTime: operatedTime,
	}

	for _, opt := range opts {
		opt(pr)
	}

	return pr
}

func (s *PRService) GetYesterdayMidnightUTCBasedOnOperatedTime() time.Time {
	return time.Date(
		s.OperatedTime.Year(), s.OperatedTime.Month(), s.OperatedTime.Day()-1,
		0, 0, 0, 0,
		time.UTC,
	)
}

func (s *PRService) GetFilteredListFromPRs(prs []models.PRItem) []models.PRItem {
	filteredPrs := []models.PRItem{}
	for _, pr := range prs {
		if pr.CreatedAt.Before(s.OperatedTime) || pr.CreatedAt.Equal(s.OperatedTime) {
			if pr.PullRequest.MergedAt != nil && pr.PullRequest.MergedAt.Before(s.GetYesterdayMidnightUTCBasedOnOperatedTime()) {
				continue
			}

			if pr.ClosedAt != nil && pr.ClosedAt.Before(s.GetYesterdayMidnightUTCBasedOnOperatedTime()) {
				continue
			}

			filteredPrs = append(filteredPrs, pr)
		}
	}

	return filteredPrs
}

func (s *PRService) GetPRState() (prState.PRState, error) {
	if s.Ent == nil {
		return prState.PRStateUnknown, fmt.Errorf("failed getting pr's state as the ent has not been yet initiate")
	}
	if s.Ent.PullRequest.MergedAt != nil && s.Ent.PullRequest.MergedAt.Before(s.OperatedTime) {
		return prState.PRStateMerged, nil
	}

	if s.Ent.ClosedAt != nil && s.Ent.ClosedAt.Before(s.OperatedTime) {
		return prState.PRStateClosed, nil
	}

	return prState.PRStateOpen, nil
}

func (s *PRService) GetPRReview(f fetcher.FetcherI) ([]models.PRReview, error) {
	if s.Ent == nil {
		return nil, fmt.Errorf("failed getting pr's review as the ent has not been yet initiate")
	}
	return f.GetPullRequestReviews(s.Ent.Number)
}
