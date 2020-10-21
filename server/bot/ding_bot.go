package bot

import (
	"github.com/CatchZeng/dingtalk"
)

func SendDingMsg(accessToken string, secret string, msg string, atMobiles []string, atAll bool) {
	bot := dingtalk.NewClient(accessToken, secret)
	dingMsg := dingtalk.NewMarkdownMessage().SetMarkdown("git message", msg).SetAt(atMobiles, atAll)
	bot.Send(dingMsg)
}
