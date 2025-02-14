package HamuFiberLib

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"test-fiber/src/global"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/gofiber/fiber/v2"
)

var DefaultFormats = map[string]huma.Format{
	"application/json": huma.DefaultJSONFormat,
	"json":             huma.DefaultJSONFormat,
}

type Route struct {
	Path        string
	ParentRoute *Route
	Tag         string
}

type RouteOptions struct {
	Tags   string
	IsAuth bool
}

func NewRouteOptions() RouteOptions {
	return RouteOptions{
		Tags:   "v1",  // Set default tags
		IsAuth: false, // Set default flags
	}
}

// Method to get the full path recursively
func (r *Route) GetFullPath() string {
	if r.ParentRoute == nil {
		return r.Path
	}
	return strings.TrimRight(r.ParentRoute.GetFullPath(), "/") + "/" + strings.TrimLeft(r.Path, "/")
}

func register[I, O any](r Route, path, method string, handler func(ctx context.Context, input *I) (*O, error), route_options RouteOptions) {

	parent_path := r.GetFullPath()

	full_path := parent_path + path

	operation := huma.Operation{
		OperationID: method + "Handler_" + full_path,
		Method:      method,
		Path:        full_path,
	}

	if route_options.IsAuth {
		operation.Security = []map[string][]string{
			{"myAuth": {}},
		}
	}

	huma.Register(global.HumaApi, operation, handler)
}

func (r *Route) Group(path string) Route {
	return Route{Path: path, ParentRoute: r}
}

func (r *Route) New(f *fiber.App, title string, version string) huma.API {
	var finalConfig huma.Config

	schemaPrefix := "#/components/schemas/"
	schemasPath := "/schemas"

	registry := huma.NewMapRegistry(schemaPrefix, huma.DefaultSchemaNamer)

	finalConfig = huma.Config{
		OpenAPI: &huma.OpenAPI{
			OpenAPI: "3.1.0",
			Info: &huma.Info{
				Title:   title,
				Version: version,
			},
			Components: &huma.Components{
				Schemas: registry,
			},
		},
		OpenAPIPath:   "/openapi",
		DocsPath:      "/docs",
		SchemasPath:   schemasPath,
		Formats:       DefaultFormats,
		DefaultFormat: "application/json",
	} // Use the default config

	finalConfig.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		// Example Authorization Code flow.
		"myAuth": {
			Type: "oauth2",
			Flows: &huma.OAuthFlows{
				AuthorizationCode: &huma.OAuthFlow{
					AuthorizationURL: "https://example.com/oauth/authorize",
					TokenURL:         "https://example.com/oauth/token",
					Scopes: map[string]string{
						"scope1": "Scope 1 description...",
						"scope2": "Scope 2 description...",
					},
				},
			},
		},

		// Example alternative describing the use of JWTs without documenting how
		// they are issued or which flows might be supported. This is simpler but
		// tells clients less information.
		"anotherAuth": {
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
		},
	}

	// Return the huma API instance with the selected config
	api := humafiber.New(f, finalConfig)
	global.HumaApi = api
	return api
}

func Post[I any, O any](r Route, path string, handler func(ctx context.Context, input *I) (*O, error), route_option RouteOptions) {

	fmt.Println("route_option:", route_option)

	register(r, path, http.MethodPost, handler, route_option)
}

func Put[I, O any](r Route, path string, handler func(ctx context.Context, input *I) (*O, error), route_option RouteOptions) {

	register(r, path, http.MethodPut, handler, route_option)
}

func Get[I, O any](r Route, path string, handler func(ctx context.Context, input *I) (*O, error), route_option RouteOptions) {

	register(r, path, http.MethodGet, handler, route_option)
}

func Delete[I, O any](r Route, path string, handler func(ctx context.Context, input *I) (*O, error), route_option RouteOptions) {

	register(r, path, http.MethodDelete, handler, route_option)
}

func CreateRoute() Route {
	return Route{}
}
