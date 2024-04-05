package events

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Luna-devv/githook/discord"
	"github.com/Luna-devv/githook/utils"
	"github.com/google/go-github/v61/github"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	var body github.PingEvent
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&body)

	discord.SendWebhook(discord.WebhookPayload{
		Username:  *body.Sender.Login,
		AvatarURL: *body.Sender.AvatarURL,
		Embeds: []discord.Embed{
			{
				Title:       fmt.Sprintf("%s: Ping", *body.Repo.FullName),
				URL:         *body.Repo.HTMLURL,
				Description: "🏓 Ping! Pong!",
				Color:       utils.GetColors().Default,
			},
		},
	})
}
