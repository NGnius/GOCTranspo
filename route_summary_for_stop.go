// NGnius 2020-01-23

package goctranspo

import (
    "encoding/json"
    "errors"
    "strconv"
)

type routeSummaryForStopInterim struct {
    StopNo string
    StopDescription string
    Error string
    Routes map[string][]simpleRouteInterim
}

func (r routeSummaryForStopInterim) Fix() RouteSummaryForStop {
    stopNo, _ := strconv.Atoi(r.StopNo)
    routes := []SimpleRoute {}
    for _, r := range r.Routes["Route"] {
        routes = append(routes, r.Fix())
    }
    return RouteSummaryForStop {
                                StopNo: stopNo,
                                StopDescription: r.StopDescription,
                                Error: r.Error,
                                Routes: routes,
                                }
}

type RouteSummaryForStop struct {
    StopNo int
    StopDescription string
    Error string
    Routes []SimpleRoute
}

func unmarshalRouteSummaryForStop(data []byte) (RouteSummaryForStop, error) {
    var rsfmMap map[string]routeSummaryForStopInterim
    var rsfm RouteSummaryForStop
    err := json.Unmarshal(data, &rsfmMap)
    if err != nil {
        rsfm.Error = err.Error()
        return rsfm, err
    }
    rsfm = rsfmMap["GetRouteSummaryForStopResult"].Fix()
    if rsfm.Error != "" {
        return rsfm, errors.New(rsfm.Error)
    }
    return rsfm, err
}
