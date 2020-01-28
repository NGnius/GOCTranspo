// NGnius 2020-01-23
package goctranspo

type simpleRouteInterim struct {
    RouteNo string
    DirectionID int
    Direction string
    RouteHeading string
}

func (r simpleRouteInterim) Fix() SimpleRoute {
    routeNo := NewRouteNo(r.RouteNo)
    return SimpleRoute {
                RouteNo: routeNo,
                DirectionID: r.DirectionID,
                Direction: r.Direction,
                RouteHeading: r.RouteHeading,
                }
}

type SimpleRoute struct {
    RouteNo RouteNo
    DirectionID int
    Direction string
    RouteHeading string
}
