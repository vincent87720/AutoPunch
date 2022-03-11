package notify

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/linebot"
)

func NotifyPunchStatus(sid string, lineID string, punchStatus string) {
	bot, err := linebot.New("channelSecret", "channelToken")
	if err != nil {
		fmt.Println(err)
	}

	msg := linebot.NewTextMessage(sid + ": " + punchStatus)
	_, err = bot.PushMessage(lineID, msg).Do()
	if err != nil {
		fmt.Println(err)
	}
}
