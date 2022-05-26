package handler

import (
	"time"

	"github.com/KaranAhlawat/ddgf/internal/app/dto"
	"github.com/KaranAhlawat/ddgf/internal/app/model"
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
		return c.Status(fiber.ErrInternalServerError.Code).
			SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(pages)
}

func (ph *PageHandler) GetPage(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString(err.Error())
	}
	page, err := ph.s.Get(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(page)
}

func (ph *PageHandler) CreatePage(c *fiber.Ctx) error {
	pageDTO := new(dto.PageHttpDTO)

	id := uuid.New()

	if err := c.BodyParser(pageDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString(err.Error())
	}

	page := model.Page{
		Datetime: time.Now(),
		Content:  pageDTO.Content,
		ID:       id,
	}

	if err := ph.s.Add(&page); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(page)
}

func (ph *PageHandler) DeletePage(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString(err.Error())
	}

	if err := ph.s.Delete(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (ph *PageHandler) UpdatePage(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString(err.Error())
	}

	pageDTO := new(dto.PageHttpDTO)
	if err = c.BodyParser(pageDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString(err.Error())
	}

	page := model.Page{
		Datetime: time.Now(),
		Content:  pageDTO.Content,
		ID:       id,
	}

	if err = ph.s.Update(id, &page); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(page)
}
