package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/irvingdinh/osgarth-api/src/component/config"
	"github.com/irvingdinh/osgarth-api/src/component/logger"
	"github.com/irvingdinh/osgarth-api/src/component/repository"
	"github.com/irvingdinh/osgarth-api/src/http/handler"
	"github.com/irvingdinh/osgarth-api/src/http/server"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start HTTP server on the predefined port",
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.CToL(context.Background(), "serveCmd")

		ctx := context.Background()

		mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(config.GetDatabaseConfig().URI))
		if err != nil {
			log.Fatalf("Failed to connect to MongoDB with error: %s", err.Error())
		}

		defer func() {
			if err := mongoClient.Disconnect(ctx); err != nil {
				log.Fatalf("Failed to disconnect from MongoDB with error: %s", err.Error())
			}
		}()

		repositoryClient := repository.New(mongoClient)

		s := server.New(
			handler.New(repositoryClient),
		)

		if err := s.Start(); err != nil {
			log.WithError(err).Fatalf("Failed to start HTTP server with error: %s", err.Error())
		}
	},
}
