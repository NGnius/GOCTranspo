// NGnius 2020-01-23

package goctranspo

import (
    "fmt"
    "testing"
    "os"
)

var (
    key string = os.Args[3]
    appID string = os.Args[4]
    debug bool = len(os.Args) > 5
)

func TestAPIGetRouteSummaryForStop(t *testing.T) {
    stopNo := 7659
    api := NewAPI(key, appID)
    rsfm, err := api.GetRouteSummaryForStop(stopNo, 0)
    if err != nil {
        t.Errorf("Error returned: %s", err)
        return
    }
    if debug {
        fmt.Println("RouteSummaryForStop")
        fmt.Println(rsfm)
    }
    // TODO test rsfm contents
}

func TestAPIGetNextTripsForStop(t *testing.T) {
    stopNo := 7659
    routeNo := 6
    api := NewAPI(key, appID)
    ntfs, err := api.GetNextTripsForStop(stopNo, routeNo)
    if err != nil {
        t.Errorf("Error returned: %s", err)
        return
    }
    if debug {
        fmt.Println("NextTripForStop")
        fmt.Println(ntfs)
    }
    if ntfs.StopNo != stopNo {
        t.Errorf("Stop was not preserved! Got %d, expected %d", ntfs.StopNo, stopNo)
    }
    if ntfs.Route.RouteNo != routeNo {
        t.Errorf("Route was not preserved! Got %d, expected %d", ntfs.Route.RouteNo, routeNo)
    }
    // TODO test ntfs contents
}

func TestAPIGetNextTripsForStopAllRoutes(t *testing.T) {
    stopNo := 7659
    api := NewAPI(key, appID)
    ntfsar, err := api.GetNextTripsForStopAllRoutes(stopNo)
    if err != nil {
        t.Errorf("Error returned: %s", err)
        return
    }
    if debug {
        fmt.Println("NextTripForStopAllRoutes")
        fmt.Println(ntfsar)
    }
    if ntfsar.StopNo != stopNo {
        t.Errorf("Stop was not preserved! Got %d, expected %d", ntfsar.StopNo, stopNo)
    }
    // TODO test ntfs contents
}
