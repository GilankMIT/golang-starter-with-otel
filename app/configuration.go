package main

import (
	"context"
	"flag"
	"github.com/allegro/bigcache"
	"go-boilerplate/common/util/bigcache_client"
	"go-boilerplate/common/util/constants"
	"go-boilerplate/common/util/logutil"
	"go-boilerplate/infra"
	"go.opentelemetry.io/otel/sdk/trace"
	"log"
	"os"
	"os/signal"
	"time"
)

func init() {

}

func StartApplication() {
	ctx := context.Background()

	ctx, tp, err := prepareTracer(ctx)
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			logutil.LogError(ctx, "failed to destroy tracer", err.Error())
			return
		}
	}()

	//client := redis.NewClient(&redis.Options{
	//	Addr:     "localhost:6379",
	//	Password: "",
	//	DB:       0,
	//})

	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Second))

	logutil.LogInfo(ctx, "starting app")

	infraOption := infra.InfraOption{
		//CacheService: redis_cache_service.NewRedisCache(client),
		CacheService: bigcache_client.NewBigCache(cache),
	}

	srv, err := infra.HandleBasicRouter(infraOption)
	if err != nil {
		logutil.LogError(ctx, "failed to start router", err.Error())
		return
	}

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

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

func prepareTracer(ctx context.Context) (context.Context, *trace.TracerProvider, error) {
	tp, err := infra.NewTraceProvider(ctx)
	if err != nil {
		logutil.LogError(ctx, "failed to initiate new application", err.Error())
		return ctx, tp, err
	}

	appTrace := tp.Tracer(constants.APP_NAME)
	ctx, _ = appTrace.Start(ctx, SERVICE_NAME)

	return ctx, tp, err
}
