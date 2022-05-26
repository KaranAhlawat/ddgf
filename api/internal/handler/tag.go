package handler

import (
	"github.com/KaranAhlawat/ddgf/internal/app/dto"
	"github.com/KaranAhlawat/ddgf/internal/app/model"
	"github.com/KaranAhlawat/ddgf/internal/app/service"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type TagHandler struct {
	s *service.Tag
}

func NewTagHandler(s *service.Tag) *TagHandler {
	return &TagHandler{
		s,
	}
}

func (th *TagHandler) GetAllTags(c *fiber.Ctx) error {
	tags, err := th.s.All()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(tags)
}

func (th *TagHandler) GetTag(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString(err.Error())
	}
	tag, err := th.s.Get(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(tag)
}

func (th *TagHandler) CreateTag(c *fiber.Ctx) error {
	tagDTO := new(dto.TagHttpDTO)

	id := uuid.New()

	if err := c.BodyParser(tagDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString(err.Error())
	}

	tag := model.Tag{
		Tag: tagDTO.Tag,
		ID:  id,
	}

	if err := th.s.Add(&tag); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(tag)
}

func (th *TagHandler) DeleteTag(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString(err.Error())
	}

	err = th.s.Delete(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (th *TagHandler) ListAllAdvices(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString(err.Error())
	}

	advices, err := th.s.ListAdvices(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(advices)
}
