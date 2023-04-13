/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/sky0621/kaubandus/cmd/setup"
	"github.com/sky0621/kaubandus/external/postgres"
	"github.com/sky0621/kaubandus/external/web"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "start api server",
	Long:  `start api server`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := setup.ReadConfig()
		setup.Logger(cfg.Env)

		db, closeDB, err := postgres.Open(cfg.Dsn(), cfg.ToDBSetOption())
		if err != nil {
			log.Err(err).Msgf("failed to setup db: %+v", err)
			return
		}
		// FIXME:
		fmt.Print(db)

		svr, shutdownServer, err := web.Server(context.Background(), cfg.Env, cfg.Trace)
		if err != nil {
			log.Err(err).Msgf("failed to setup server: %+v", err)
			return
		}
		defer shutdownServer()

		go func() {
			q := make(chan os.Signal, 1)
			signal.Notify(q, os.Interrupt, syscall.SIGTERM)
			<-q
			closeDB()
			shutdownServer()
			os.Exit(-1)
		}()

		if err := svr.ListenAndServe(":" + cfg.WebPort); err != nil {
			log.Err(err).Msgf("failed to start server: %+v", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
