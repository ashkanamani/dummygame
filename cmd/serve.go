package cmd

import (
	"github.com/ashkanamani/dummygame/internal/repository"
	"github.com/ashkanamani/dummygame/internal/repository/redis"
	"github.com/ashkanamani/dummygame/internal/service"
	"github.com/ashkanamani/dummygame/internal/telegram"
	"github.com/sirupsen/logrus"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "running the bot",
	Run:   serve,
}

func serve(_ *cobra.Command, _ []string) {

	// Setup repositories
	redisClient, err := redis.NewRedisClient(os.Getenv("REDIS_URL"))
	if err != nil {
		logrus.WithError(err).Fatal("could not connect to redis server.")
	}
	// set up app
	accountRepository := repository.NewAccountRedisRepository(redisClient)
	accountService := service.NewAccountService(accountRepository)
	app := service.NewApp(accountService)

	tg, err := telegram.NewTelegram(app, os.Getenv("TELEGRAM_API_TOKEN"))
	if err != nil {
		slog.Error("could not connect to telegram servers.", "error", err)
		os.Exit(1)
	}
	slog.Info("Starting the Bot.")
	tg.Start()

}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
