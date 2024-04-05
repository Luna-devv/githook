package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Luna-devv/githook/config"
)

var client = &http.Client{}

func SendWebhook(payload WebhookPayload) {
	conf := config.Get()
	payloadBytes, err := json.Marshal(payload)

	if err != nil {
		fmt.Println("Error occurred", err)
		return
	}

	fmt.Println(bytes.NewBuffer(payloadBytes))
	req, err := http.NewRequest("POST", conf.Webhook, bytes.NewBuffer(payloadBytes))

	if err != nil {
		fmt.Println("Error occurred", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Error occurred", err)
		return
	}

	if res.StatusCode != http.StatusNoContent {
		fmt.Println("Error occurred", res.Status)
		body := new(bytes.Buffer)
		body.ReadFrom(res.Body)
		fmt.Println(body.String())
	}

	res.Body.Close()
}
