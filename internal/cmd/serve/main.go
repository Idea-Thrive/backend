package serve

import (
	"github.com/Idea-Thrive/backend/internal/config"
	"github.com/Idea-Thrive/backend/internal/http/handler"
	"github.com/Idea-Thrive/backend/internal/logger"
	"github.com/Idea-Thrive/backend/internal/store"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "run http server",
		Long:  "run http server",
		Run:   main,
	}

	return cmd
}

func main(cmd *cobra.Command, args []string) {
	cfg := config.Load("config.yaml")

	logger := logger.New(cfg.Log)

	s := store.NewStore(nil)

	app := fiber.New()

	handler.Authentication{
		Secret: cfg.Http.Secret,
		Logger: logger,
		Store:  s,
	}.Register(app)

	log.Fatal(app.Listen(":" + strconv.Itoa(cfg.Http.Port)))
}
