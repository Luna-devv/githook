package events

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Luna-devv/githook/discord"
	"github.com/Luna-devv/githook/utils"
	"github.com/google/go-github/v61/github"
)

func Push(w http.ResponseWriter, r *http.Request, url string) {
	var body github.PushEvent
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&body)

	if len(body.Commits) == 0 {
		return
	}

	var desc string
	for _, c := range body.Commits {
		commit := *c
		desc += fmt.Sprintf(
			"[`%s`](%s) %s\n",
			(*commit.ID)[:7],
			*commit.URL, utils.Truncate(*commit.Message, 62),
		)
	}

	branch := strings.Split(*body.Ref, "/heads/")[1]

	discord.SendWebhook(
		url,
		discord.WebhookPayload{
			Username:  *body.Sender.Login,
			AvatarURL: *body.Sender.AvatarURL,
			Embeds: []discord.Embed{
				{
					Title: fmt.Sprintf(
						"%s%s: %d commit%s",
						*body.Repo.FullName,
						utils.Ternary(
							branch == "" || branch == "master" || branch == "main",
							"",
							"@"+branch,
						),
						len(body.Commits),
						utils.Ternary(len(body.Commits) > 1, "s", ""),
					),
					URL:         *body.Compare,
					Description: utils.Truncate(desc, 4000),
					Color:       utils.GetColors().Default,
				},
			},
		},
	)
}
