package services

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/Thiti-Dev/gh-pr-gamify-v1/models"
	"github.com/Thiti-Dev/gh-pr-gamify-v1/pkg/slack"
	prState "github.com/Thiti-Dev/gh-pr-gamify-v1/types/pr-state"
	"github.com/Thiti-Dev/gh-pr-gamify-v1/utils"
)

type PRSummaryConcluderService struct {
	summarizationDate time.Time
	slack             slack.SlackI
	staging           []PRSummaryConcluderItem
}

type PRSummaryConcluderItem struct {
	FormattedRepositoryPath string
	CreatedAt               time.Time
	Status                  prState.PRState
	Ent                     models.PRItem
	Reviewers               []models.PRReview
}

func NewPRSummaryConcluderService(summarizationDate time.Time, staging []PRSummaryConcluderItem, slack slack.SlackI) *PRSummaryConcluderService {
	return &PRSummaryConcluderService{
		summarizationDate: summarizationDate,
		staging:           staging,
		slack:             slack,
	}
}

func (s *PRSummaryConcluderService) GetTotalStagingConcluderItem() int {
	return len(s.staging)
}

// Sort sorts the staging slice by FormattedRepositoryPath (ascending),
// and within each repository path, by CreatedAt (ascending).
func (svc *PRSummaryConcluderService) Sort() []PRSummaryConcluderItem {

	filteredList := []PRSummaryConcluderItem{}

	filteredList = append(filteredList, svc.staging...)

	sort.Slice(filteredList, func(i, j int) bool {
		// Primary sort by FormattedRepositoryPath
		if filteredList[i].FormattedRepositoryPath != filteredList[j].FormattedRepositoryPath {
			return filteredList[i].FormattedRepositoryPath < filteredList[j].FormattedRepositoryPath
		}

		// Secondary: Sort by state (descending)
		if filteredList[i].Status != filteredList[j].Status {
			// Assuming a lexicographical string sort and reversing order for descending
			return filteredList[i].Status > filteredList[j].Status
		}

		// Tertiary sort by CreatedAt
		return filteredList[i].CreatedAt.Before(filteredList[j].CreatedAt)
	})

	return filteredList
}

func (svc *PRSummaryConcluderService) SummarizeIntoSlackChannel() error {
	yesterdayMidnightUTC := time.Date(
		svc.summarizationDate.Year(), svc.summarizationDate.Month(), svc.summarizationDate.Day()-1,
		0, 0, 0, 0,
		time.UTC,
	)
	formattedYesterdayMidnightUTC := yesterdayMidnightUTC.Format("2006-01-02 03:04PM")
	formattedOperateDate := svc.summarizationDate.Format("2006-01-02 03:04PM")

	var formattedOnGoingPR string
	onGoingPRs := svc.Sort()

	formattedHeader := fmt.Sprintf("%-40s %-8s %-50s", "Title", "Status", "Approver")
	var currentRepositoryPath string

	mappedApprovingCountByLogin := map[string]int{}

	for i, pr := range onGoingPRs {
		fmt.Println(pr.FormattedRepositoryPath)

		// Check if repository path has changed or if we're on the first entry
		if i == 0 || currentRepositoryPath != pr.FormattedRepositoryPath {
			// Close the previous repository block if it's not the first item
			if i > 0 {
				formattedOnGoingPR += "```"
			}

			// Start a new repository block
			currentRepositoryPath = pr.FormattedRepositoryPath
			formattedOnGoingPR += fmt.Sprintf("\nrepository: `%s`\n```\n%s\n%s\n", currentRepositoryPath, formattedHeader, strings.Repeat("-", 32*2))
		}

		// TODO: Segregate this into its place
		formattedApprovers := "none"
		approvers := []string{}
		for _, review := range pr.Reviewers {
			if review.State == "APPROVED" && review.SubmittedAt.Before(svc.summarizationDate) {
				mappedApprovingCountByLogin[review.User.Login]++
				approvers = append(approvers, review.User.Login)
			}
		}

		if len(approvers) != 0 {
			formattedApprovers = strings.Join(approvers, ", ")
		}

		// Append the PR information (title, status, approver)
		formattedOnGoingPR += fmt.Sprintf("%-40s %-8v %-50s\n", utils.TruncateString(pr.Ent.Title, 40), pr.Status, formattedApprovers)

		// If it's the last PR and we are on the last repository, close the block
		if i == len(onGoingPRs)-1 && currentRepositoryPath == pr.FormattedRepositoryPath {
			formattedOnGoingPR += "```"
		}
	}

	var formattedApprovingScore string

	if len(mappedApprovingCountByLogin) != 0 {
		formattedApprovingScore = fmt.Sprintf("Gamification result 🎮:\n```%-10s %-30s %-5s %-5s\n%s\n", "Place", "Name", "Score", "Earn", strings.Repeat("-", 32*2))
		sortedApprovingCount := svc.createSortedApproveCountFromMap(mappedApprovingCountByLogin)

		for i, ac := range sortedApprovingCount {
			renderedMedal := "🧩"
			if i == 0 {
				renderedMedal = "🥇"
			} else if i == 1 {
				renderedMedal = "🥈"
			} else if i == 2 {
				renderedMedal = "🥉"
			}

			formattedApprovingScore += fmt.Sprintf("%-10d %-30s %-5d %-5s\n", i+1, ac.Key, ac.Value, renderedMedal)
		}
		formattedApprovingScore += "```"
	}

	err := svc.slack.Post(fmt.Sprintf("Github PR Notify/Summarization (%s) - (%s)\n%s\n\n%s", formattedYesterdayMidnightUTC, formattedOperateDate, formattedOnGoingPR, formattedApprovingScore))

	return err
}

func (svc *PRSummaryConcluderService) createSortedApproveCountFromMap(m map[string]int) []models.KV {

	var sortedSlice []models.KV
	for k, v := range m {
		sortedSlice = append(sortedSlice, models.KV{Key: k, Value: v})
	}

	sort.Slice(sortedSlice, func(i, j int) bool {
		return sortedSlice[i].Value > sortedSlice[j].Value
	})

	return sortedSlice
}
