package main

import (
	log "github.com/sirupsen/logrus"
	"log/slog"
)

func main() {
	slog.Info("Starting server.")
	log.SetFormatter(
		&log.TextFormatter{
			ForceColors: true,
		},
	)
}
