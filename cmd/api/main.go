package main

import (
	"fmt"
	"itemmeli/package/cache"
	"itemmeli/package/config"
	"itemmeli/package/server"
	"itemmeli/package/service"

	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

var (
	GitCommit string
	GitTag    string
	BuildTime string
)

func main() {
	cli.VersionPrinter = func(cCtx *cli.Context) {
		fmt.Printf("Git Tag: %s\n", GitTag)
		fmt.Printf("Git Commit: %s\n", GitCommit)
		fmt.Printf("Build Time: %s\n", BuildTime)
	}
	app := &cli.App{
		Name:            "MercadoLibre API",
		Version:         GitTag,
		HideHelpCommand: true,
		HideVersion:     false,
		Description:     "MercadoLibre API for item details page",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
			},
		},
		Action: runServer,
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal().Msg(err.Error())
	}
}

func runServer(cCtx *cli.Context) error {
	var (
		err        error
		configYaml *config.YamlConfig
	)

	signalCtx, cancel := signal.NotifyContext(cCtx.Context, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	configYaml, err = config.NewYamlConfig(cCtx.String("config"))
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	cache := cache.NewRedisCache(configYaml)
	service := service.NewService(cache)
	server := server.NewServerV1(service, configYaml)

	go func() {
		log.Info().Msgf("Starting API server: %s \n", server.Info())
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("API server failed: %v", err)
		}

	}()

	<-signalCtx.Done() // waiting for signal to stop the server
	log.Info().Msg("Shut down data server")
	if err = server.Stop(cCtx.Context); err != nil {
		log.Fatal().Msg("Can't terminate data server")
	}

	return nil
}
