package main

import (
	"context"
	"flag"
	"go-otel/app/infra"
	"go-otel/app/util/constants"
	"go-otel/app/util/logutil"
	"log"
	"os"
	"os/signal"
	"time"
)

const SERVICE_NAME = "main"

func main() {
	ctx := context.Background()
	tp, err := infra.NewTraceProvider(ctx)

	appTrace := tp.Tracer(constants.APP_NAME)
	ctx, _ = appTrace.Start(ctx, SERVICE_NAME)
	if err != nil {
		logutil.LogError(ctx, "failed to initiate new application", err.Error())
		return
	}

	logutil.LogInfo(ctx, "starting app")

	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			logutil.LogError(ctx, "failed to destroy tracer", err.Error())
			return
		}
	}()

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	srv, err := infra.HandleBasicRouter()
	if err != nil {
		logutil.LogError(ctx, "failed to start router", err.Error())
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
