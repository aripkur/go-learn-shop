package auth

import "github.com/gofiber/fiber/v2"

type handler struct {
	svc service
}

func newHandler(svc service) handler {
	return handler{
		svc: svc,
	}
}

func (h handler) register(ctx *fiber.Ctx) error {
	var req = RegisterRequestPayload{}
}
