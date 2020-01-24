// NGnius 2020-01-23
package goctranspo

import (
    "strconv"
)

type simpleRouteInterim struct {
    RouteNo string
    DirectionID int
    Direction string
    RouteHeading string
}

func (r simpleRouteInterim) Fix() SimpleRoute {
    routeNo, _ := strconv.Atoi(r.RouteNo)
    return SimpleRoute {
                RouteNo: routeNo,
                DirectionID: r.DirectionID,
                Direction: r.Direction,
                RouteHeading: r.RouteHeading,
                }
}

type SimpleRoute struct {
    RouteNo int
    DirectionID int
    Direction string
    RouteHeading string
}
