package runner

import (
	"fmt"
	"time"

	"github.com/Thiti-Dev/gh-pr-gamify-v1/core/fetcher"
	"github.com/Thiti-Dev/gh-pr-gamify-v1/models"
	"github.com/Thiti-Dev/gh-pr-gamify-v1/pkg/config"
	"github.com/Thiti-Dev/gh-pr-gamify-v1/pkg/slack"
	"github.com/Thiti-Dev/gh-pr-gamify-v1/services"
)

type Runner struct {
	githubFetchers []fetcher.FetcherI
	appConfig      *config.Config
}

func NewRunner(fetchProcessList []fetcher.FetcherI, conf *config.Config) *Runner {
	return &Runner{
		githubFetchers: fetchProcessList,
		appConfig:      conf,
	}
}

func (r *Runner) Run() error {
	slackAlert := slack.NewSlack(r.appConfig.SlackWebhookURL)

	now := time.Now().UTC()
	nDaysAgo := now.AddDate(0, 0, -365)

	basedPRServ := services.NewPRService(now)
	summaryCacheControl := services.NewPRSummaryCacheControl()

	stagingSummarization := make([]services.PRSummaryConcluderItem, 0)

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
		summarizedList, err := summaryServ.Summarize()
		if err != nil {
			return err
		}

		stagingSummarization = append(stagingSummarization, summarizedList...)
	}

	summaryConcluder := services.NewPRSummaryConcluderService(now, stagingSummarization, slackAlert)

	err := summaryConcluder.SummarizeIntoSlackChannel()

	return err
}
