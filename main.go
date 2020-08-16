package main

import (
	"context"
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
		val, err := rdb.Incr(context.Background(), "counter").Result()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write([]byte(fmt.Sprintf("Hello world! x%d", val)))
	})

	port := os.Getenv("HTTP_PORT")

	log.Printf("listening on port %s\n", port)

	panic(http.ListenAndServe(":"+port, nil))
}
