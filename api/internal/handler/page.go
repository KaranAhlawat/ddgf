package handler

import (
	"fmt"
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
		c.Status(fiber.ErrInternalServerError.Code).SendString(err.Error())
		return err
	}
	return c.JSON(pages)
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
	return c.JSON(page)
}

func (ph *PageHandler) CreatePage(c *fiber.Ctx) error {
	pageDTO := new(dto.PagePostDTO)

	if err := c.BodyParser(pageDTO); err != nil {
		c.Status(fiber.StatusBadGateway).
			SendString(err.Error())
		return err
	}

	page := model.Page{
		Datetime: time.Now(),
		Content:  pageDTO.Content,
		ID:       pageDTO.ID,
	}

	if err := ph.s.Add(&page); err != nil {
		c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
		return err
	}

	return c.JSON(page)
}

func (ph *PageHandler) DeletePage(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		c.Status(fiber.StatusBadRequest).
			SendString(fmt.Errorf("invalid UUIDv4 : %w", err).Error())
		return err
	}

	if err := ph.s.Delete(id); err != nil {
		c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
		return err
	}

	return nil
}

func (ph *PageHandler) UpdatePage(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		c.Status(fiber.StatusBadRequest).
			SendString(fmt.Errorf("invalid UUIDv4 : %w", err).Error())
		return err
	}

	pageDTO := new(dto.PagePostDTO)
	if err := c.BodyParser(pageDTO); err != nil {
		c.Status(fiber.StatusBadGateway).
			SendString(err.Error())
		return err
	}

	page := model.Page{
		Datetime: time.Now(),
		Content:  pageDTO.Content,
		ID:       id,
	}

	if err := ph.s.Update(id, &page); err != nil {
		c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
		return err
	}

	return c.JSON(page)
}
