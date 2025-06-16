package events

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/google/go-github/v61/github"
	"github.com/redis/go-redis/v9"
)

type WorkflowJobCache struct {
	ID          string `json:"id"`
	RunID       string `json:"run_id"`
	Name        string `json:"name"`
	Conclusions string `json:"conclusions"`
}

func WorkflowJob(w http.ResponseWriter, r *http.Request, client *redis.Client) {
	var body github.WorkflowJobEvent
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&body)

	if *body.Action != "completed" {
		return
	}

	data := WorkflowJobCache{
		ID:          strconv.FormatInt(*body.WorkflowJob.ID, 10),
		RunID:       strconv.FormatInt(*body.WorkflowJob.RunID, 10),
		Name:        *body.WorkflowJob.Name,
		Conclusions: *body.WorkflowJob.Conclusion,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = client.Set(
		r.Context(),
		fmt.Sprintf("workflow:job:%s:%d", *body.WorkflowJob.HeadSHA, *body.WorkflowJob.ID),
		jsonData,
		24*time.Hour,
	).
		Err()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
