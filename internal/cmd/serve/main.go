package serve

import (
	"log"
	"strconv"

	"github.com/Idea-Thrive/backend/internal/config"
	"github.com/Idea-Thrive/backend/internal/http/handler"
	"github.com/Idea-Thrive/backend/internal/http/middleware"
	"github.com/Idea-Thrive/backend/internal/jwt"
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
	cfg := config.Load("config.yaml")

	logger := logger.New(cfg.Log)

	str := store.NewStore(nil)

	app := fiber.New()

	j := jwt.NewJWT(cfg.JWT)

	auth := middleware.Auth{
		JWT:    j,
		Logger: logger,
		Store:  str,
	}

	handler.Authentication{
		JWT:    j,
		Logger: logger,
		Store:  str,
	}.Register(app.Group("/auth"))

	handler.User{
		Store:  str,
		Logger: logger,
	}.Register(app.Group("/users", auth.Auth))

	handler.Idea{
		Store:  str,
		Logger: logger,
	}.Register(app.Group("/ideas", auth.Auth))

	log.Fatal(app.Listen(":" + strconv.Itoa(cfg.HTTP.Port)))
}
