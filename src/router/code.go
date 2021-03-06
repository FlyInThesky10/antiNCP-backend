package router

import (
	"github.com/FlyInThesky10/antiNCP-backend/controller"
	"github.com/FlyInThesky10/antiNCP-backend/middleware"
	"github.com/labstack/echo/v4"
)

func InitCodeRouter(g *echo.Group) {
	g.Use(middleware.JWTMiddleware())
	g.PUT("/submitCode", controller.UserSubmitCode)
	g.GET("/viewSubmission", controller.UserViewSubmission)
	g.GET("/allSubmission", controller.AdminGetSubmission)
	g.PUT("/verifySubmission", controller.AdminVerifySubmission)
}
