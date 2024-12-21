package services

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Thiti-Dev/gh-pr-gamify-v1/core/fetcher"
	"github.com/Thiti-Dev/gh-pr-gamify-v1/models"
)

type PRSummaryService struct {
	OperatedTime          time.Time
	PRs                   []PRService
	CacheCtrl             *PRSummaryCacheControl
	Fetcher               fetcher.FetcherI
	IsReadyToBeSummarized bool
}

type PRSummaryCacheControl struct {
	mu                              sync.RWMutex
	CachedReviewByPullRequestNumber map[int][]models.PRReview
}

// Set sets a value in the map safely.
func (c *PRSummaryCacheControl) Set(key int, value []models.PRReview) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.CachedReviewByPullRequestNumber[key] = value
}

// Get safely reads from the map.
func (c *PRSummaryCacheControl) Get(key int) ([]models.PRReview, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, exists := c.CachedReviewByPullRequestNumber[key]
	return val, exists
}

func NewPRSummaryCacheControl() *PRSummaryCacheControl {
	return &PRSummaryCacheControl{
		CachedReviewByPullRequestNumber: make(map[int][]models.PRReview),
	}
}

func NewPRSummaryService(t time.Time, prs []PRService, prcctrl *PRSummaryCacheControl, f fetcher.FetcherI) *PRSummaryService {
	return &PRSummaryService{
		OperatedTime: t,
		PRs:          prs,
		CacheCtrl:    prcctrl,
		Fetcher:      f,
	}
}

func (s *PRSummaryService) CollectReviews() {
	var wg sync.WaitGroup

	for _, prs := range s.PRs {
		if _, found := s.CacheCtrl.Get(prs.Ent.Number); found {
			continue // prevent collected review to be fetched again
		}

		wg.Add(1)
		go func() {
			defer wg.Done()

			review, err := prs.GetPRReview(s.Fetcher)
			if err != nil {
				log.Fatalf("failed getting PR-Review: %d, err: %s", prs.Ent.Number, err)
			}

			s.CacheCtrl.Set(prs.Ent.Number, review)
		}()
	}

	wg.Wait()

	s.IsReadyToBeSummarized = true
}

func (s *PRSummaryService) Summarize() error {
	if !s.IsReadyToBeSummarized {
		return fmt.Errorf("not yet been ready for summarizing")
	}

	return nil
}
