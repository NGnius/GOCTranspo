// NGnius 2020-01-23
package goctranspo

import (
    "strconv"
    "time"
)

type routeDirectionInterim struct {
    RouteNo string
    RouteLabel string
    Direction string
    Error string
    RequestProcessingTime string
    Trips map[string][]tripInterim
}

func (r routeDirectionInterim) Fix() RouteDirection {
    routeNo, _ := strconv.Atoi(r.RouteNo)
    requestProcessingTime, _ := time.Parse("20060102150405", r.RequestProcessingTime) // YYYYMMDDHHmmSS (HH in 24-hour format)
    trips := []Trip {}
    for _, t := range r.Trips["Trip"] {
        trips = append(trips, t.Fix())
    }
    return RouteDirection {
                RouteNo: routeNo,
                RouteLabel: r.RouteLabel,
                Direction: r.Direction,
                Error: r.Error,
                RequestProcessingTime: requestProcessingTime,
                Trips: trips,
                }
}

type RouteDirection struct {
    RouteNo int
    RouteLabel string
    Direction string
    Error string
    RequestProcessingTime time.Time
    Trips []Trip
}
