package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
	"os"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		visitors, err := rdb.Incr(r.Context(), "counter").Result()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write([]byte(fmt.Sprintf("Hello world! You are visitor #%d", visitors)))
	})

	port := os.Getenv("HTTP_PORT")

	log.Printf("listening on port %s\n", port)

	panic(http.ListenAndServe(":"+port, nil))
}
