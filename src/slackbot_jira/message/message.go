package message

import (
	"fmt"
	"slackbot_jira/atlassian"
	"slackbot_jira/config"
)

type Message struct {
	SlackChannel string
	AsUser       config.SlackUser
	Text         string
}

type MessageMatcher interface {
	GetMatchingMessages([]config.MessageTrigger, ...atlassian.ActivityIssue) []Message
}

type matcher struct {
	cfg config.SlackConfig
}

func NewMessageMatcher(cfg config.SlackConfig) MessageMatcher {
	return matcher{cfg}
}

func (m matcher) GetMatchingMessages(triggers []config.MessageTrigger, activity_issues ...atlassian.ActivityIssue) []Message {
	messages := make([]Message, 0)

	for _, activity_issue := range activity_issues {
		for _, trigger := range triggers {
			if match, ok := m.get_match(trigger, activity_issue); ok {
				messages = append(messages, match.get_messages()...)
			}
		}
	}

	return messages
}

func (m matcher) get_match(trigger config.MessageTrigger, activity_issue atlassian.ActivityIssue) (match, bool) {
	// TODO
	return match{m.cfg.Users, trigger, activity_issue}, true
}

type match struct {
	users          map[string]config.SlackUser
	trigger        config.MessageTrigger
	activity_issue atlassian.ActivityIssue
}

func (m match) get_messages() []Message {
	message := Message{
		m.trigger.SlackChannel,
		m.users[m.trigger.SlackUserKey],
		fmt.Sprintf("%s", m.activity_issue.Activity.Title),
	}
	return []Message{message}
}
