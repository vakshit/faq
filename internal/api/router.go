package api

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/questions/all", getAllQuestions)
	app.Get("/questions/approved", getApprovedQuestions)
	app.Post("/questions/post", postQuestion)
	app.Post("/questions/answer", answerQuestion)
	app.Delete("/questions/delete", deleteQuestion)
}
