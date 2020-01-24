// NGnius 2020-01-23

package goctranspo

import (
    "encoding/json"
    "errors"
    "strconv"
)

type nextTripsForStopInterim struct {
    StopNo string
    StopLabel string
    Error string
    Route map[string]routeDirectionInterim
}

func (n nextTripsForStopInterim) Fix() NextTripsForStop {
    stopNo, _ := strconv.Atoi(n.StopNo)
    route := n.Route["RouteDirection"].Fix()
    return NextTripsForStop {
                            StopNo: stopNo,
                            StopLabel: n.StopLabel,
                            Error: n.Error,
                            Route: route,
                            }
}

type NextTripsForStop struct {
    StopNo int
    StopLabel string
    Error string
    Route RouteDirection
}

func unmarshalNextTripsForStop(data []byte) (NextTripsForStop, error) {
    var ntfsMap map[string]nextTripsForStopInterim
    var ntfs NextTripsForStop
    err := json.Unmarshal(data, &ntfsMap)
    if err != nil {
        ntfs.Error = err.Error()
        return ntfs, err
    }
    ntfs = ntfsMap["GetNextTripsForStopResult"].Fix()
    if ntfs.Error != "" {
        return ntfs, errors.New(ntfs.Error)
    }
    return ntfs, err
}
