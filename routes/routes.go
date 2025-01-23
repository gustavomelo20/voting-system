package routes

import (
	"voting-system/controllers"
	"voting-system/repositories"
	"voting-system/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	repo := repositories.NewCandidateRepository(db)
	service := services.NewCandidateService(repo)
	controller := controllers.NewCandidateController(service)

	e.GET("/candidates", controller.GetCandidates)
	e.POST("/candidate/:id/vote", controller.Vote)
	e.GET("/candidate/:id/vote", controller.GetCandidateVotes)
}
