// NGnius 2020-01-23

package goctranspo

import (
    "fmt"
    "testing"
    "os"
)

var (
    key string = os.Getenv("GOCTRANSPO_KEY")
    appID string = os.Getenv("GOCTRANSPO_APPID")
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
    if ntfs.Route.RouteNo.Number != routeNo {
        t.Errorf("Route was not preserved! Got %d, expected %d", ntfs.Route.RouteNo.Number, routeNo)
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

func TestGTFSAgency(t *testing.T) {
    table := "routes"
    id := 34
    column := "route_short_name"
    value := "42" // BUG Spaces (and presumably other HTML-escaped chars) cause a 400 error
    order_by := "id"
    direction := "ASC" // ASC or DESC only (or empty)
    limit := 10
    api := NewAPI(key, appID)
    gtfs, err := api.GTFS(table, id, column, value, order_by, direction, limit)
    if err != nil {
        t.Errorf("Error returned: %s", err)
        return
    }
    if debug {
        fmt.Println("GTFS")
        fmt.Printf("Processed %d results\n", len(gtfs.Gtfs))
        fmt.Println(gtfs.Query)
        fmt.Println(gtfs.Gtfs)
    }
    if gtfs.Query.Table != table {
        t.Errorf("Table was not preserved! Got %s, expected %s", gtfs.Query.Table, table)
    }
    if column != "" && gtfs.Query.Column != column {
        t.Errorf("Column was not preserved! Got %s, expected %s", gtfs.Query.Column, column)
    }
    if value != "" && gtfs.Query.Value != value {
        t.Errorf("Value was not preserved! Got %s, expected %s", gtfs.Query.Value, value)
    }
    if order_by != "" && gtfs.Query.Order_By != order_by {
        t.Errorf("Order_By was not preserved! Got %s, expected %s", gtfs.Query.Order_By, order_by)
    }
    if limit != 0 && gtfs.Query.Limit != limit {
        t.Errorf("Limit was not preserved! Got %d, expected %d", gtfs.Query.Limit, limit)
    }
    if gtfs.Query.Format != "json" {
        t.Errorf("Format was not preserved! Got %s, expected %s", gtfs.Query.Format, "json")
    }
    // TODO test gtfs contents
}
