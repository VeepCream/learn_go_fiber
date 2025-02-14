package main

import (
	"fmt"
	"test-fiber/src/routes"
	hamu_fiber_lib "test-fiber/src/utils/HamuFiberLib"

	"github.com/danielgtaylor/huma/v2"
	"github.com/gofiber/fiber/v2"
)

func MyMiddleware(ctx huma.Context, next func(huma.Context)) {
	test := ctx.Operation().Security

	fmt.Println("test:", test)
	if len(test) > 0 {
		fmt.Println("test[0]:", test[0])
	}
	next(ctx)
}

func main() {
	// Create a new router & API
	app := fiber.New()
	route := hamu_fiber_lib.CreateRoute()
	api := route.New(app, "My API", "v1.0.0")
	api.UseMiddleware(MyMiddleware)
	routes.Routes(route)

	app.Listen(":8080")
}
