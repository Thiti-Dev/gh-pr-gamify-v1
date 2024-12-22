package slack

import (
	"github.com/slack-go/slack"
)

type SlackI interface {
	Post(message string) error
}

type Slack struct {
	webhook string
}

func NewSlack(webhook string) SlackI {
	return &Slack{
		webhook: webhook,
	}
}

func (s *Slack) Post(message string) error {
	msg := slack.WebhookMessage{
		Text: message,
	}

	err := slack.PostWebhook(s.webhook, &msg)
	if err != nil {
		return err
	}

	return nil
}
