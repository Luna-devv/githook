package main

import (
	"fmt"
	"net/http"

	"github.com/Luna-devv/githook/config"
	"github.com/Luna-devv/githook/events"
	_ "github.com/joho/godotenv/autoload"
	"github.com/redis/go-redis/v9"
)

var conf = config.Get()

var client = redis.NewClient(&redis.Options{
	Addr:     conf.Redis.Addr,
	Password: conf.Redis.Password,
	Username: conf.Redis.Username,
	DB:       conf.Redis.Db,
})

func main() {
	http.HandleFunc("POST /", handler)

	http.ListenAndServe(":8080", nil)

	defer client.Close()
}

func handler(w http.ResponseWriter, r *http.Request) {
	githubEvent := r.Header.Get("X-Github-Event")

	if githubEvent == "" {
		http.Error(w, "Missing X-Github-Event header", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received GitHub event: %s\n", githubEvent)

	switch githubEvent {
	case "ping":
		events.Ping(w, r)
	case "push":
		events.Push(w, r)
	default:
		http.Error(w, fmt.Sprintf("Unsupported event: %s", githubEvent), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
