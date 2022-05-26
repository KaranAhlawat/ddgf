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
		return c.Status(200).JSON(&fiber.Map{
			"ok": true,
		})
	})

	api := s.app.Group("/api/", logger.New())
	pageAPI := api.Group("/page/")
	adviceAPI := api.Group("/advice/")
	tagAPI := api.Group("/tag/")

	ph := handler.NewPageHandler(ps)
	ah := handler.NewAdviceHandler(as)
	th := handler.NewTagHandler(ts)

	pageAPI.Get("/", ph.GetAllPages)
	pageAPI.Get("/:id/", ph.GetPage)
	pageAPI.Post("/", ph.CreatePage)
	pageAPI.Delete("/:id/", ph.DeletePage)
	pageAPI.Put("/:id", ph.UpdatePage)

	adviceAPI.Get("/", ah.GetAdvices)
	adviceAPI.Get("/:id/", ah.GetAdvice)
	adviceAPI.Post("/", ah.CreateAdvice)
	adviceAPI.Delete("/:id/", ah.DeleteAdvice)
	adviceAPI.Put("/:id/", ah.UpdateAdvice)
	adviceAPI.Get("/tag/:aid/:tid/", ah.Tag)
	adviceAPI.Get("/untag/:aid/:tid/", ah.Untag)

	tagAPI.Get("/", th.GetAllTags)
	tagAPI.Get("/:id/", th.GetTag)
	tagAPI.Post("/", th.CreateTag)
	tagAPI.Delete("/:id", th.DeleteTag)
	tagAPI.Get("/:id/advices/", th.ListAllAdvices)
}
