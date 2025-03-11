package handlers

import (
	"coworking/internal/adapters/http/models"
	"coworking/internal/core/usecases/commands"
	"coworking/internal/ports"

	"github.com/gofiber/fiber/v2"
)

type MembershipHandler struct {
	createMembershipUseCase *commands.CreateMembershipUseCase
}

func NewMembershipHandler(createUseCase *commands.CreateMembershipUseCase) *MembershipHandler {
	return &MembershipHandler{createMembershipUseCase: createUseCase}
}

func (h *MembershipHandler) RegisterRoutes(app *fiber.App) {
	app.Group("/memberships").Post("/", h.CreateMembership)
}

func (h *MembershipHandler) CreateMembership(c *fiber.Ctx) error {
	req, err := parseAndValidateMembershipRequest(c)
	if err != nil {
		return err
	}

	params := commands.CreateMembershipParams{
		UserID: req.UserID,
	}

	if err := h.createMembershipUseCase.Handle(params); err != nil {
		return handleDomainError(c, err)
	}

	return c.SendStatus(fiber.StatusCreated)
}

func parseAndValidateMembershipRequest(c *fiber.Ctx) (*models.MembershipDTO, error) {
	var req models.MembershipDTO

	if err := c.BodyParser(&req); err != nil {
		return nil, FormatErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	if validationErrors := req.Validate(); len(validationErrors) > 0 {
		return nil, FormatErrorResponse(c, fiber.StatusBadRequest, "Validation failed", validationErrors)
	}

	return &req, nil
}

var _ ports.HttpPort = (*MembershipHandler)(nil)
