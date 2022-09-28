package cmd

import (
	"github.com/vakshit/faq/internal/api"
	"github.com/vakshit/faq/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the Fiber Server",
	RunE: func(cmd *cobra.Command, args []string) error {
		database.ConnectMongo()
		err := Serve()
		return err
	},
}

func Serve() error {
	// Create a new fiber instance with custom config
	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's an fiber.*Error
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			// Send custom error page
			err = ctx.Status(code).JSON(fiber.Map{"message": err.Error()})
			if err != nil {
				// In case the SendFile fails
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal Server Error"})
			}

			// Return from handler
			return nil
		},
	})

	api.SetupRoutes(app)
	err := app.Listen(":5000")
	return err
}
