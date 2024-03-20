package app

import (
	"diploma/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/hashicorp/go-hclog"
)

type Server struct {
	App    *fiber.App
	Logger hclog.Logger
}

func Start(conf *config.Config) error {
	s := new(Server)
	s.App = fiber.New(fiber.Config{
		BodyLimit: 20 * 1024 * 1024,
	})

	s.App.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "OPTIONS, GET, POST, HEAD, PUT, DELETE, PATCH",
		AllowHeaders:     "Origin,X-Requested-With, Content-Type, Accept, Authorization",
		AllowCredentials: true,
		ExposeHeaders:    "",
		MaxAge:           120,
	}))

	s.Router()

	err := s.App.Listen(":" + conf.Port)

	return err
}
