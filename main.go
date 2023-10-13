package main

import (
	"context"
	"flag"
	infra2 "go-otel/app/infra"
	"go-otel/app/util/constants"
	log_util2 "go-otel/app/util/log_util"
	"log"
	"os"
	"os/signal"
	"time"
)

const SERVICE_NAME = "main"

func main() {
	ctx := context.Background()
	tp, err := infra2.NewTraceProvider(ctx)

	appTrace := tp.Tracer(constants.SERVICE_NAME)
	ctx, _ = appTrace.Start(ctx, SERVICE_NAME)
	if err != nil {
		log_util2.LogError(ctx, "failed to initiate new application", err.Error())
		return
	}

	log_util2.LogInfo(ctx, "starting app")

	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log_util2.LogError(ctx, "failed to destroy tracer", err.Error())
			return
		}
	}()

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	srv, err := infra2.HandleBasicRouter()
	if err != nil {
		log_util2.LogError(ctx, "failed to start router", err.Error())
		return
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	<-ctx.Done()
	log.Println("shutting down")
	os.Exit(0)
}
