package ru

import (
	"context"
	"fmt"
	hamu_fiber_lib "test-fiber/src/utils/HamuFiberLib"

	"github.com/danielgtaylor/huma/v2"
	"github.com/gofiber/fiber/v2"
)

type GreetRequest struct {
	Name string `json:"name" validate:"required"`
}

type GreetResponse struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

func PostHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "This is v1/ru/post",
	})
}

type GreetingInput struct {
	Body struct {
		KKb string
	}
}

type PathInput struct {
	Name  string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
	Tag   string `query:"tag" enum:"foo,bar,baz"`
	Sales uint   `query:"sales" maximum:"1000"`
}

type GreetingOutput struct {
	MyHeader string `json:"myHeader"`
	Body     struct {
		KKt string `json:"message"`
	}
}

func PostRoute(route hamu_fiber_lib.Route) {

	hamu_fiber_lib.Get(route, "/1/{name}", func(ctx context.Context, input *PathInput) (*GreetingOutput, error) {
		if input.Tag == "bob" {
			return nil, huma.Error404NotFound("no greeting for bob")
		}
		resp := &GreetingOutput{}
		resp.MyHeader = "MyValue"
		resp.Body.KKt = fmt.Sprintf("Hello, %s!", input.Name)
		return resp, nil
	}, hamu_fiber_lib.RouteOptions{Tags: "v1", IsAuth: true})

	hamu_fiber_lib.Put(route, "/2", func(ctx context.Context, input *GreetingInput) (*GreetingOutput, error) {
		if input.Body.KKb == "bob" {
			return nil, huma.Error404NotFound("no greeting for bob")
		}
		resp := &GreetingOutput{}
		resp.MyHeader = "MyValue"
		resp.Body.KKt = fmt.Sprintf("Hello, %s!", input.Body.KKb)
		return resp, nil
	}, hamu_fiber_lib.RouteOptions{Tags: "v1"})

	hamu_fiber_lib.Post(route, "/3", func(ctx context.Context, input *GreetingInput) (*GreetingOutput, error) {
		if input.Body.KKb == "bob" {
			return nil, huma.Error404NotFound("no greeting for bob")
		}
		resp := &GreetingOutput{}
		resp.MyHeader = "MyValue"
		resp.Body.KKt = fmt.Sprintf("Hello, %s!", input.Body.KKb)
		return resp, nil
	}, hamu_fiber_lib.RouteOptions{Tags: "v1", IsAuth: true})

	hamu_fiber_lib.Delete(route, "/4", func(ctx context.Context, input *GreetingInput) (*GreetingOutput, error) {
		if input.Body.KKb == "bob" {
			return nil, huma.Error404NotFound("no greeting for bob")
		}
		resp := &GreetingOutput{}
		resp.MyHeader = "MyValue"
		resp.Body.KKt = fmt.Sprintf("Hello, %s!", input.Body.KKb)
		return resp, nil
	}, hamu_fiber_lib.RouteOptions{Tags: "v1", IsAuth: true})

}
