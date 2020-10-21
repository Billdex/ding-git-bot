package controller

import (
	"ding-git-bot/model/webhook"
	"fmt"
	"strings"
)

func makePushMessage(repositoryName string, branch string, url string, sender string, commits []webhook.Commit) string {
	text := fmt.Sprintf("[%s:%s](%s) %s 提交了 %d 个commit \n", repositoryName, branch, url, sender, len(commits))
	for _, commit := range commits {
		text += fmt.Sprintf("> [%s](%s) %s - %s \n",
			commit.Id[0:8], commit.Url,
			strings.Replace(commit.Message, "\n", " ", -1),
			commit.Author.Name)
	}
	return text
}
