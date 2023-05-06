/*
Copyright © 2023 sky0621 <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/sky0621/familiagildo/app"
	log2 "github.com/sky0621/familiagildo/app/log"
	"github.com/sky0621/familiagildo/cmd/setup"
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
		cfg := app.ReadConfig()
		log2.Logger(cfg.Env)

		app, err := setup.InitializeApp(cfg)
		if err != nil {
			log.Err(err).Msg("failed to initialize app")
			return
		}
		defer app.Close()

		go func() {
			q := make(chan os.Signal, 1)
			signal.Notify(q, os.Interrupt, syscall.SIGTERM)
			<-q
			defer app.Close()
			os.Exit(-1)
		}()

		if err := app.Start(cfg.WebPort); err != nil {
			log.Err(err).Msg("failed to start app")
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
