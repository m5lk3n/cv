// Package main is the cvapi binary, a small HTTP server exposing the CV
// content as a JSON API.
//
//	@title			CV API
//	@version		1.0
//	@description	HTTP API exposing CV content (about, basics, work, etc.) as JSON.
//	@BasePath		/
package main

import (
	"flag"
	"log/slog"
	"os"

	"lttl.dev/cv/api"
	_ "lttl.dev/cv/docs"
)

func main() {
	addr := flag.String("addr", ":8080", "address to listen on")
	flag.Parse()

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	slog.Info("starting cvapi", "addr", *addr)
	if err := api.Run(*addr); err != nil {
		slog.Error("server failed", "err", err)
		os.Exit(1)
	}
}
