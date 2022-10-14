package api

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/all", getAllQuestions)
	app.Get("/approved", getApprovedQuestions)
	app.Post("/post", postQuestion)
	app.Post("/answer", answerQuestion)
	app.Delete("/delete", deleteQuestion)
}
