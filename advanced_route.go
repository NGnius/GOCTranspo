// NGnius 2020-01-23
package goctranspo

type advancedRouteInterim struct {
    RouteNo string
    DirectionID int
    Direction string
    RouteHeading string
    Trips []tripInterim
}

func (r advancedRouteInterim) Fix() AdvancedRoute {
    routeNo := NewRouteNo(r.RouteNo)
    trips := []Trip {}
    for _, t := range r.Trips {
        trips = append(trips, t.Fix())
    }
    return AdvancedRoute {
                RouteNo: routeNo,
                DirectionID: r.DirectionID,
                Direction: r.Direction,
                RouteHeading: r.RouteHeading,
                Trips: trips,
                }
}

type AdvancedRoute struct {
    RouteNo RouteNo
    DirectionID int
    Direction string
    RouteHeading string
    Trips []Trip
}
