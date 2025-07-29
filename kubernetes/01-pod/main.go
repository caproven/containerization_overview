package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	instance := os.Getenv("INSTANCE")
	if instance == "" {
		instance = "default instance"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Received request",
			slog.String("path", r.URL.Path),
			slog.String("method", r.Method),
			slog.String("host", r.RemoteAddr),
		)
		data := fmt.Sprintf("hello from %s", instance)
		_, _ = w.Write([]byte(data))
	})

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}
	slog.Info("Listening", slog.String("addr", addr))
	if err := http.ListenAndServe(addr, nil); err != nil {
		slog.Error("HTTP server exited", "err", err)
	}
}
