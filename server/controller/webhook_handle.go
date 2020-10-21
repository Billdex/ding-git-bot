package controller

import (
	"ding-git-bot/bot"
	"ding-git-bot/log"
	"ding-git-bot/model/database"
	"ding-git-bot/model/webhook"
	"ding-git-bot/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

// 解析WebHook消息来源于事件类型
func WebHookParsing(c *gin.Context) {
	// 判断消息来源并解析事件类型
	if event := c.GetHeader("X-GitHub-Event"); event != "" {
		// 消息来源为GitHub
		// 解析Payload并获取仓库地址
		payload, err := c.GetRawData()
		if err != nil {
			log.Error("获取Payload失败!", err)
			return
		}
		msg := webhook.GithubMsg{}
		err = json.Unmarshal(payload, &msg)
		if err != nil {
			log.Error("解析Payload为Struct失败!", err)
			return
		}
		repository := new(database.Repository)
		has, err := database.DBEngine.Where("url = ?", msg.Repository.HtmlUrl).Get(repository)
		if err != nil {
			log.Error("查询数据库出错!", err)
			return
		}
		if !has {
			log.Warn("仓库地址未记录:", msg.Repository.HtmlUrl)
			return
		}

		// 验证签名,signature为空时不需要验证
		signature := c.GetHeader("X-Hub-Signature")
		if signature != "" {
			if correct := utils.VerifySignature(string(payload), signature, repository.Secret); !correct {
				log.Warn("签名有误！仓库:", msg.Repository.HtmlUrl, "签名:", signature)
				return
			}
		}

		// 处理消息事件
		githubEventHandle(payload, event)
		return
	} else if event := c.GetHeader("X-Gitlab-Event"); event != "" {
		// 消息来源为GitLab

		return
	} else if event := c.GetHeader("X-Gogs-Event"); event != "" {
		// 消息来源为Gogs
		// 解析Payload并获取仓库地址
		payload, err := c.GetRawData()
		if err != nil {
			log.Error("获取Payload失败!", err)
			return
		}
		signature := c.GetHeader("X-Gogs-Signature")
		// 处理消息事件
		gogsEventHandle(payload, event, signature)

		return
	} else {
		// 未知来源
		data, _ := c.GetRawData()
		log.Info("WebHook信息来源不明：", data)
		return
	}
}

// 处理github消息事件
func githubEventHandle(payload []byte, event string) {
	switch event {
	case EventPing:

	case EventStar:

	case EventPush:

	default:
		log.Warn("未知消息类型:", event)
	}
	return
}

//处理gogs消息事件
func gogsEventHandle(payload []byte, event string, signature string) {
	msg := webhook.GogsMsg{}
	err := json.Unmarshal(payload, &msg)
	if err != nil {
		log.Error("解析Payload为Struct失败!", err)
		return
	}
	repository := new(database.Repository)
	has, err := database.DBEngine.Where("url = ?", msg.Repository.HtmlUrl).Get(repository)
	if err != nil {
		log.Error("查询数据库出错!", err)
		return
	}
	if !has {
		log.Warn("仓库地址未记录:", msg.Repository.HtmlUrl)
		return
	}

	// 验证签名,signature为空时不需要验证
	if signature != "" {
		if correct := utils.VerifySignature(string(payload), signature, repository.Secret); !correct {
			log.Warn("签名有误！仓库:", msg.Repository.HtmlUrl, "签名:", signature)
			return
		}
	}

	var baseMsg string
	switch event {
	case EventPush:
		baseMsg, err = gogsPushHandle(payload)
	default:
		log.Warn("未知消息类型:", event)
		return
	}
	if err != nil {
		log.Error("解析Payload为Struct失败!", err)
		return
	}

	notices := make([]database.ViewNoticeEvent, 0)
	err = database.DBEngine.Where("repository_id = ? and event = ?", repository.Id, EventPush).Find(&notices)
	if err != nil {
		log.Error("查询数据库出错!", err)
		return
	}

	if len(notices) == 0 {
		log.Infof("仓库%s未设置Push事件通知!", msg.Repository.Name)
		return
	}

	for _, notice := range notices {
		if notice.NoticeForbidden {
			log.Infof("Notice %d 已禁用!", notice.NoticeId)
			continue
		}

		senders := make([]database.ViewNoticeSender, 0)
		err = database.DBEngine.Where("notice_id = ?", notice.NoticeId).Find(&senders)
		if err != nil {
			//查询出错则跳过该notice的处理，不影响其他notice信息处理
			log.Error("查询数据库出错!", err)
			continue
		}

		//若notice未设定任何sender，则视为处理所有sender的消息
		//若notice设定了sender，则其它未包含的sender的消息都跳过
		if len(senders) != 0 {
			log.Infof("Notice %d 指定了 %d 个sender", notice.NoticeId, len(senders))
			hasSender := false
			for _, sender := range senders {
				if sender.SenderEmail == msg.Sender.Email {
					hasSender = true
					break
				}
			}
			if !hasSender {
				log.Infof("消息来源不属于指定sender，不推送该Notice")
				continue
			}
		}

		followers := make([]database.ViewNoticeFollower, 0)
		err = database.DBEngine.Where("notice_id = ?", notice.NoticeId).Find(&followers)
		if err != nil {
			//查询出错则跳过该notice的处理，不影响其他notice信息处理
			log.Error("查询数据库出错!", err)
			continue
		}
		dingMsg := baseMsg + "\n" + notice.Message
		atMobile := make([]string, 0)
		for _, follower := range followers {
			atMobile = append(atMobile, follower.FollowerPhone)
			dingMsg += "@" + follower.FollowerPhone
		}

		//通过bot发送钉钉消息
		bots := make([]database.ViewNoticeBot, 0)
		err = database.DBEngine.Where("notice_id = ?", notice.NoticeId).Find(&bots)
		if err != nil {
			//查询出错则跳过该notice的处理，不影响其他notice信息处理
			log.Error("查询数据库出错!", err)
			continue
		}
		for _, msgBot := range bots {
			if msgBot.BotForbidden {
				log.Infof("bot %d 已禁用！", msgBot.BotId)
				continue
			}
			bot.SendDingMsg(msgBot.AccessToken, msgBot.Secret, dingMsg, atMobile, msgBot.AtAll)
			log.Info("发送了一条推送到", msgBot.AccessToken, "\n内容为", baseMsg+"\n"+notice.Message)
		}
	}
}
