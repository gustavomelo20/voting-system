package controllers

import (
	"net/http"
	"strconv"
	"voting-system/services"

	"github.com/labstack/echo/v4"
)

type CandidateController struct {
	service services.CandidateService
}

func NewCandidateController(service services.CandidateService) *CandidateController {
	return &CandidateController{service: service}
}

func (c *CandidateController) GetCandidates(ctx echo.Context) error {
	candidates, err := c.service.GetAllCandidates()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "Erro ao buscar candidatos"})
	}
	return ctx.JSON(http.StatusOK, candidates)
}

func (c *CandidateController) Vote(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "ID inválido"})
	}

	err = c.service.VoteForCandidate(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "Voto registrado com sucesso"})
}

func (c *CandidateController) GetCandidateVotes(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "ID inválido"})
	}

	uintID := uint(id)

	votes, err := c.service.GetCandidateVotes(uintID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "Erro ao buscar votos"})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"votes": votes})

}
