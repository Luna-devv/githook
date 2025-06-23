package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var client = &http.Client{}

func SendWebhook(url string, payload WebhookPayload) {
	payloadBytes, err := json.Marshal(payload)

	if err != nil {
		fmt.Println("Error occurred", err)
		return
	}

	log.Println(bytes.NewBuffer(payloadBytes))
	req, err := http.NewRequest("POST", url + "?with_components=true", bytes.NewBuffer(payloadBytes))

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
		log.Fatalln("Error occurred", res.Status)
		body := new(bytes.Buffer)
		body.ReadFrom(res.Body)
		fmt.Println(body.String())
	}

	res.Body.Close()
}
