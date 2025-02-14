package routes

import (
	v1 "test-fiber/src/routes/v1"
	hamu_fiber_lib "test-fiber/src/utils/HamuFiberLib"
)

func Routes(route hamu_fiber_lib.Route) {
	routev1 := route.Group("/v1")

	v1.RouteV(routev1)
}
