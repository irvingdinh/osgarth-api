package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/irvingdinh/osgarth-api/src/component/logger"
	"github.com/irvingdinh/osgarth-api/src/http/handler"
	"github.com/irvingdinh/osgarth-api/src/http/server"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start HTTP server on the predefined port",
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.CToL(context.Background(), "serveCmd")

		s := server.New(
			handler.New(),
		)

		if err := s.Start(); err != nil {
			log.WithError(err).Fatalf("Failed to start HTTP server with error: %s", err.Error())
		}
	},
}
