package v1

import (
	"test-fiber/src/routes/v1/ru"
	hamu_fiber_lib "test-fiber/src/utils/HamuFiberLib"
)

func RouteV(route hamu_fiber_lib.Route) {
	// ru.RuRoute(app, router, api)
	// aa := app.GetRoutes()
	// for _, route := range aa {
	// 	fmt.Printf("Hello, %s!\n", route.Path)
	// }

	route_ru := route.Group("/ru")
	ru.PostRoute(route_ru)
}
