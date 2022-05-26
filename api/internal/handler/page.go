package handler

import (
	"fmt"

	"github.com/KaranAhlawat/ddgf/internal/app/service"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PageHandler struct {
	s *service.Page
}

func NewPageHandler(s *service.Page) *PageHandler {
	return &PageHandler{
		s,
	}
}

func (ph *PageHandler) GetAllPages(c *fiber.Ctx) error {
	pages, err := ph.s.All()
	if err != nil {
		c.Status(fiber.ErrInternalServerError.Code).SendString(err.Error())
		return err
	}
	err = c.JSON(pages)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		return err
	}
	return nil
}

func (ph *PageHandler) GetPage(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		c.Status(fiber.StatusBadRequest).
			SendString(fmt.Errorf("invalid UUIDv4 : %w", err).Error())
		return err
	}
	page, err := ph.s.Get(id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
		return err
	}
	err = c.JSON(page)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
		return err
	}
	return nil
}
