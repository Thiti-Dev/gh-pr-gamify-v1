package runner

import (
	"fmt"
	"time"

	"github.com/Thiti-Dev/gh-pr-gamify-v1/core/fetcher"
	"github.com/Thiti-Dev/gh-pr-gamify-v1/models"
	"github.com/Thiti-Dev/gh-pr-gamify-v1/services"
)

type Runner struct {
	githubFetchers []fetcher.FetcherI
}

func NewRunner(fetchProcessList []fetcher.FetcherI) *Runner {
	return &Runner{
		githubFetchers: fetchProcessList,
	}
}

func (r *Runner) Run() error {
	_now := time.Now().UTC()

	yesterdayMidnightUTC := time.Date(
		_now.Year(), _now.Month(), _now.Day()-11,
		0, 0, 0, 0,
		time.UTC,
	)

	now := yesterdayMidnightUTC

	nDaysAgo := now.AddDate(0, 0, -365)

	basedPRServ := services.NewPRService(now)
	summaryCacheControl := services.NewPRSummaryCacheControl()

	// TODO: iterate over the fetchers
	for _, fetcher := range r.githubFetchers {
		res, err := fetcher.GetPullRequestList(models.DateRange{From: nDaysAgo, To: now})
		if err != nil {
			return err
		}
		fmt.Println(res.TotalCount)

		filteredList := basedPRServ.GetFilteredListFromPRs(res.Items)
		fmt.Println(len(filteredList))

		prServs := []services.PRService{}

		for _, f := range filteredList {
			prServ := services.NewPRService(now, services.WithPRItem(f))
			prServs = append(prServs, *prServ)
			prState, _ := prServ.GetPRState()
			fmt.Println(f.Title, prState)
		}

		summaryServ := services.NewPRSummaryService(now, prServs, summaryCacheControl, fetcher)
		summaryServ.CollectReviews()
		err = summaryServ.Summarize()
		if err != nil {
			return err
		}
	}
	return nil
}
