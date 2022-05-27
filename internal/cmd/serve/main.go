package serve

import (
	"github.com/Idea-Thrive/backend/internal/mysql"
	"github.com/Idea-Thrive/backend/internal/store/operation"
	"go.uber.org/zap"
	"log"
	"strconv"

	"github.com/Idea-Thrive/backend/internal/config"
	"github.com/Idea-Thrive/backend/internal/http/handler"
	"github.com/Idea-Thrive/backend/internal/logger"
	"github.com/Idea-Thrive/backend/internal/store"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

// Command function.
func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "run http server",
		Long:  "run http server",
		Run:   main,
	}

	return cmd
}

// main function.
func main(cmd *cobra.Command, args []string) {
	cfg := config.Load("config.yml")

	logger := logger.New(cfg.Log)

	db, err := mysql.New(cfg.DB, logger)
	if err != nil {
		logger.Fatal("error in database", zap.Error(err))
	}
	connection := operation.NewOperation(db, logger)

	str := store.NewStore(connection)

	app := fiber.New()

	handler.Authentication{
		Secret: cfg.HTTP.Secret,
		Logger: logger,
		Store:  str,
	}.Register(app.Group("/auth"))

	handler.User{
		Store:  str,
		Logger: logger,
	}.Register(app.Group("/users"))

	log.Fatal(app.Listen(":" + strconv.Itoa(cfg.HTTP.Port)))
}
