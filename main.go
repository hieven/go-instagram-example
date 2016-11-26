package main

import (
	instagram "github.com/hieven/go-instagram"
	"github.com/robfig/cron"
)

const InstagramBotID = 0
const Username = ""
const Password = ""

func main() {
	ch := make(chan int)

	ig, _ := instagram.Create(Username, Password)

	ig.Login()

	c := cron.New()

	c.AddFunc("*/5 * * * * *", func() {
		ig.Inbox.GetFeed()

		if ig.Inbox.Threads[0].Items[0].UserID != InstagramBotID {
			switch ig.Inbox.Threads[0].Items[0].ItemType {
			case "placeholder":
				ig.Inbox.Threads[0].BroadcastText("Oops, It's a private message.")
			case "media_share":
				ig.Inbox.Threads[0].BroadcastText("Thanks for sharing this media to me.")
			case "location":
				ig.Inbox.Threads[0].BroadcastText("Amazing location!")
			case "text":
				ig.Inbox.Threads[0].BroadcastText("Hello, How are you?")
			case "like":
				ig.Inbox.Threads[0].BroadcastText("I love u, too")
			case "hashtag":
				ig.Inbox.Threads[0].BroadcastText("Is it a trending hashtag?")
			case "media":
				ig.Inbox.Threads[0].BroadcastText("Wonderful selfie :)")
			default:
				return
			}
		}
	})

	c.Start()
	<-ch
}
