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

	ctx := r.Context()

	data := WorkflowJobCache{
		ID:          strconv.FormatInt(*body.WorkflowJob.ID, 10),
		Name:        *body.WorkflowJob.Name,
		Conclusions: *body.WorkflowJob.Conclusion,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = client.Set(
		ctx,
		fmt.Sprintf("workflow:%d", *body.WorkflowJob.ID),
		jsonData,
		24*time.Hour,
	).
		Err()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
