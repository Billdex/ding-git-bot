package controller

import (
	"ding-git-bot/model/webhook"
	"encoding/json"
	"strings"
)

//处理gogs的push事件消息
func gogsPushHandle(payload []byte) (string, error) {
	msg := webhook.GogsPushMsg{}
	err := json.Unmarshal(payload, &msg)
	if err != nil {
		return "", err
	}
	ref := strings.Split(msg.Ref, "/")
	branch := ref[len(ref)-1]
	commits := make([]webhook.Commit, 0)
	for _, gogsCommit := range msg.Commits {
		commits = append(commits, webhook.Commit{
			Id:      gogsCommit.Id,
			Message: gogsCommit.Message,
			Url:     gogsCommit.Url,
			Author: webhook.Author{
				Name:  gogsCommit.Author.Name,
				Email: gogsCommit.Author.Email,
			},
		})
	}
	baseMsg := makePushMessage(msg.Repository.Name, branch, msg.CompareUrl, msg.Sender.Username, commits)
	return baseMsg, nil
}
