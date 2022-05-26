package handler

import (
	"github.com/KaranAhlawat/ddgf/internal/app/dto"
	"github.com/KaranAhlawat/ddgf/internal/app/model"
	"github.com/KaranAhlawat/ddgf/internal/app/service"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AdviceHandler struct {
	s *service.Advice
}

func NewAdviceHandler(s *service.Advice) *AdviceHandler {
	return &AdviceHandler{
		s,
	}
}

func (ah *AdviceHandler) GetAdvices(c *fiber.Ctx) error {
	advices, err := ah.s.All()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(advices)
}

func (ah *AdviceHandler) GetAdvice(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString(err.Error())
	}

	advice, err := ah.s.Get(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(advice)
}

func (ah *AdviceHandler) CreateAdvice(c *fiber.Ctx) error {
	adviceDTO := new(dto.AdviceHttpDTO)

	if err := c.BodyParser(adviceDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString(err.Error())
	}

	advice := model.Advice{
		Content: adviceDTO.Content,
		Tags:    []model.Tag{},
		ID:      uuid.New(),
	}

	err := ah.s.Add(&advice)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(advice)
}

func (ah *AdviceHandler) DeleteAdvice(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString(err.Error())
	}

	err = ah.s.Delete(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (ah *AdviceHandler) UpdateAdvice(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString(err.Error())
	}

	adviceDTO := new(dto.AdviceHttpDTO)
	if err = c.BodyParser(adviceDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString(err.Error())
	}

	advice := model.Advice{
		Content: adviceDTO.Content,
		ID:      id,
		Tags:    []model.Tag{},
	}

	if err = ah.s.Update(id, &advice); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(advice)
}

func (ah *AdviceHandler) Tag(c *fiber.Ctx) error {
	aid, err := uuid.Parse(c.Params("aid"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString(err.Error())
	}

	tid, err := uuid.Parse(c.Params("tid"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString(err.Error())
	}

	err = ah.s.AddTag(aid, tid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}

func (ah *AdviceHandler) Untag(c *fiber.Ctx) error {
	aid, err := uuid.Parse(c.Params("aid"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString(err.Error())
	}

	tid, err := uuid.Parse(c.Params("tid"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString(err.Error())
	}

	err = ah.s.Untag(aid, tid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}
