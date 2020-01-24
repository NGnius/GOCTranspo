// NGnius 2020-01-23

package goctranspo

import (
    "encoding/json"
    "errors"
    "strconv"
)

type nextTripsForStopAllRoutesInterim struct {
    StopNo string
    StopDescription string
    Error string
    Routes map[string][]advancedRouteInterim
}

func (n nextTripsForStopAllRoutesInterim) Fix() NextTripsForStopAllRoutes {
    stopNo, _ := strconv.Atoi(n.StopNo)
    routes := []AdvancedRoute {}
    for _, r := range n.Routes["Route"] {
        routes = append(routes, r.Fix())
    }
    return NextTripsForStopAllRoutes {
                                StopNo: stopNo,
                                StopDescription: n.StopDescription,
                                Error: n.Error,
                                Routes: routes,
                                }
}

type NextTripsForStopAllRoutes struct {
    StopNo int
    StopDescription string
    Error string
    Routes []AdvancedRoute
}

func unmarshalNextTripsForStopAllRoutes(data []byte) (NextTripsForStopAllRoutes, error) {
    var ntfsarMap map[string]nextTripsForStopAllRoutesInterim
    var ntfsar NextTripsForStopAllRoutes
    err := json.Unmarshal(data, &ntfsarMap)
    if err != nil {
        ntfsar.Error = err.Error()
        return ntfsar, err
    }
    ntfsar = ntfsarMap["GetRouteSummaryForStopResult"].Fix()
    if ntfsar.Error != "" {
        return ntfsar, errors.New(ntfsar.Error)
    }
    return ntfsar, err
}
