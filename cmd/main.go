package main

import (
	"fmt"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	ytd "github.com/kkdai/youtube/v2"

	emoji "github.com/kyokomi/emoji/v2"
)

func main() {
	token := "5187223070:AAGnIEUbBzwalGrVR3eE5yd9VFAMXDqQR2U"
	bot, err := tgbot.NewBotAPI(token)

	if err != nil {
		panic(err)
	}

	bot.Debug = true

	updateConfig := tgbot.NewUpdate(0)

	updateConfig.Timeout = 20

	updates := bot.GetUpdatesChan(updateConfig)

	for v := range updates {

		if v.Message == nil {
			continue
		}

		msg := tgbot.NewMessage(v.Message.Chat.ID, v.Message.Text+emoji.Sprintf(":pizza:, :snake:"))

		// msg.ReplyToMessageID = v.Message.MessageID

		// res, err := http.Get(v.Message.Text)
		// if err != nil {
		// 	fmt.Println(err)
		// }

		video := ytd.Client{Debug: true}

		// text, err := io.ReadAll(res.Body)
		// if err != nil {
		// 	panic(err)
		// }

		videoStr, err := video.GetVideo(v.Message.Text)
		if err != nil {
			fmt.Println(err)
		}

		// fmt.Println("\n start \n\n\n",videoStr.Formats,"\n\n end\n")

		for _,v := range videoStr.Formats {
			// if v.QualityLabel == "720p" {
			// 	fmt.Println("\n\n\n",v,"\n\n\n")
			// }
			fmt.Print( v.QualityLabel)

			
		}

		// fmt.Println(string(text))

		_, err = bot.Send(msg)
		// if err != nil {
		// 	fmt.Println(err)
		// }

		// fmt.Println(v.Message.Text)

	}
}
