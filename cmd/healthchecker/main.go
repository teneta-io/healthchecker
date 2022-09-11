package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/teneta-io/healthchecker/pkg/apiserver"
	"github.com/teneta-io/healthchecker/pkg/config"
	"github.com/teneta-io/healthchecker/pkg/keystore"
	"github.com/teneta-io/healthchecker/pkg/logger"
	"github.com/teneta-io/healthchecker/pkg/stresser"
)

func main() {
	config := config.NewConfig()
	logger.NewLogger(logger.LogLevel(config.Log.Level), config.Log.Mode == "dev")

	keyStore, err := keystore.NewKeyStore()
	if err != nil {
		logger.Fatalf("Can`t load keys")
	}

	stresser := stresser.NewStresser()
	go stresser.Run()

	apiServer := apiserver.NewServer(":9999", &config.Server, keyStore, stresser)

	logger.Infof("Initialization done.")

	runErr := make(chan error, 1)
	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.Infof("HTTP API server start listen on: %s", apiServer.Addr)
		if err := apiServer.ListenAndServe(); err != nil {
			runErr <- fmt.Errorf("can't start http server: %w", err)
		}
	}()

	select {
	case err := <-runErr:
		logger.Fatalf("Running error: %s", err)
	case s := <-quitCh:
		logger.Infof("Received signal: %v. Running graceful shutdown...", s)

		ctx, done := context.WithTimeout(context.Background(), config.Server.ShutdownTimeout.Duration)
		defer done()

		if err := apiServer.Shutdown(ctx); err != nil {
			logger.Infof("Can't shutdown API server: %s", err)
		}
	}
}
