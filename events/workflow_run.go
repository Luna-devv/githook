package events

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Luna-devv/githook/discord"
	"github.com/Luna-devv/githook/utils"
	"github.com/google/go-github/v61/github"
	"github.com/redis/go-redis/v9"
)

func WorkflowRun(w http.ResponseWriter, r *http.Request, url string, client *redis.Client) {
	var body github.WorkflowRunEvent
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&body)

	if *body.Action != "completed" {
		return
	}

	ctx := r.Context()
	jobKeys := client.Keys(ctx, "workflow:*")
	defer client.Del(ctx, jobKeys.Val()...)
	var desc string

	for _, key := range jobKeys.Val() {
		data, err := client.Get(ctx, key).Result()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var job WorkflowJobCache
		err = json.Unmarshal([]byte(data), &job)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		desc += fmt.Sprintf(
			"%s %s [↗︎](%s)\n",
			utils.Ternary(job.Conclusions == "success", "<:tick:1017781086102761543>", "<:cross:1017781065340964934>"),
			job.Name,
			fmt.Sprintf("https://github.com/%s/actions/runs/%d/job/%s", *body.Repo.FullName, *body.WorkflowRun.ID, job.ID),
		)
	}

	head := ""
	if body.WorkflowRun.HeadBranch != nil {
		head = *body.WorkflowRun.HeadBranch
	}

	discord.SendWebhook(
		url,
		discord.WebhookPayload{
			Username:  *body.Sender.Login,
			AvatarURL: *body.Sender.AvatarURL,
			Embeds: []discord.Embed{
				{
					Title: fmt.Sprintf(
						"%s%s: Workflow %s",
						*body.Repo.FullName,
						utils.Ternary(
							head == "" || head == "master" || head == "main",
							"",
							"@"+head,
						),
						*body.WorkflowRun.Conclusion,
					),
					Description: fmt.Sprintf(
						"[`%s`](%s) %s",
						(*body.WorkflowRun.HeadCommit.ID)[:7],
						fmt.Sprintf("https://github.com/%s/commit/%s", *body.Repo.FullName, *body.WorkflowRun.HeadCommit.ID),
						utils.Truncate(*body.WorkflowRun.HeadCommit.Message, 62),
					) + "\n\n>>> " + desc,
					URL: *body.WorkflowRun.HTMLURL,
					Color: utils.Ternary(
						*body.WorkflowRun.Conclusion == "success",
						utils.GetColors().Success,
						utils.GetColors().Error,
					).(int),
				},
			},
		},
	)
}
