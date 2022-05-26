package server

import (
	"github.com/KaranAhlawat/ddgf/internal/app/service"
	"github.com/KaranAhlawat/ddgf/internal/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	app  fiber.App
	conf APIConfig
}

func NewServer(apiConf *APIConfig) *Server {
	return &Server{
		app:  *fiber.New(),
		conf: *NewAPIConfig(),
	}
}

func (s *Server) Run() {
	s.app.Listen(s.conf.port)
}

func (s *Server) Setup(ps *service.Page, as *service.Advice, ts *service.Tag) {
	s.app.Get("/healthcheck", func(c *fiber.Ctx) error {
		c.Status(200).JSON(&fiber.Map{
			"ok": true,
		})
		return nil
	})

	api := s.app.Group("/api/", logger.New())

	ph := handler.NewPageHandler(ps)

	api.Get("/page/", ph.GetAllPages)
	api.Get("/page/:id", ph.GetPage)
}
