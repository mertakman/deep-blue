package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mertakman/deep-blue/src/pathgenerator"
)

func CalculateRoute(ctx *fiber.Ctx) error {
	startSquare := ctx.Query("start_square")
	dstSquare := ctx.Query("dst_square")

	steps, err := pathgenerator.GenerateKnightPath(startSquare, dstSquare)
	if err != nil {
		// At least one of given coordinates are in wrong format , or at outside of matrix.
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}

	resp := new(CalculatedRouteResp)
	resp.Steps = steps

	// returning the provided response to
	return ctx.Status(http.StatusOK).JSON(resp)
}

type CalculatedRouteResp struct {
	Steps []string `json:"steps"` // contains the steps between 2 squares as algebric annotations.
}
