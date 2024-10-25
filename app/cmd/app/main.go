package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Alexander272/si_bot/internal/config"
	"github.com/Alexander272/si_bot/internal/server"
	"github.com/Alexander272/si_bot/internal/services"
	transport "github.com/Alexander272/si_bot/internal/transport/http"
	"github.com/Alexander272/si_bot/pkg/logger"
	"github.com/Alexander272/si_bot/pkg/mattermost"
)

func main() {
	// if err := gotenv.Load("../.env"); err != nil {
	// 	logger.Fatalf("failed to load env variables. error: %s", err.Error())
	// }

	conf, err := config.Init("configs/config.yaml")
	if err != nil {
		logger.Fatalf("failed to init configs. error: %s", err.Error())
	}
	logger.Init(os.Stdout, conf.Environment)

	//* Dependencies
	mattermostConf := mattermost.Config{
		ServerLink: conf.Most.ServerLink,
		Token:      conf.Most.Token,
	}
	mostClient := mattermost.NewMattermostClient(mattermostConf)

	_, _, err = mostClient.Http.GetPing()
	if err != nil {
		logger.Fatalf("failed to ping most. error: %s", err.Error())
	}

	//* Services, Repos & API Handlers
	servicesDeps := services.Deps{
		MostClient: mostClient.Http,
		BotName:    conf.Most.BotName,
	}
	services := services.NewServices(servicesDeps)
	handlers := transport.NewHandler(services)

	//* HTTP Server
	srv := server.NewServer(conf, handlers.Init(conf))
	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Fatalf("error occurred while running http server: %s\n", err.Error())
		}
	}()
	logger.Infof("Application started on port: %s", conf.Http.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second
	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server. error: %s", err.Error())
	}
}
