// NGnius 2020-01-23

package goctranspo

import (
    "fmt"
    "net/http"
    "net/url"
    "strconv"
    "io/ioutil"
    "os"
)

const (
    urlGetRouteSummaryForStop = "https://api.octranspo1.com/v1.3/GetRouteSummaryForStop"
    urlGetNextTripsForStop = "https://api.octranspo1.com/v1.3/GetNextTripsForStop"
    urlGetNextTripsForStopAllRoutes = "https://api.octranspo1.com/v1.3/GetNextTripsForStopAllRoutes"
    urlGTFS = "https://api.octranspo1.com/v1.3/Gtfs"
)

var (
    debug bool = os.Getenv("GOCTRANSPO_DEBUG") != ""
)

type API struct {
    Key string
    AppID string
    client *http.Client
}

func NewAPI(key string, appID string) *API {
    newAPI := &API {
                    Key: key,
                    AppID: appID,
                    client: &http.Client {},
                    }
    return newAPI
}

func (api *API) startOfBody() string {
    return "appID="+api.AppID+"&apiKey="+api.Key+"&format=json"
}

func (api *API) startOfForm() url.Values {
    v := url.Values{}
    v.Add("appID", api.AppID)
    v.Add("apiKey", api.Key)
    v.Add("format", "json")
    return v
}

func (api *API) GetRouteSummaryForStop(stopNo int, routeNo int) (RouteSummaryForStop, error) {
    var rsfm RouteSummaryForStop
    values := api.startOfForm()
    values.Add("stopNo", strconv.Itoa(stopNo))
    if routeNo != 0 {
        values.Add("routeNo", strconv.Itoa(routeNo))
    }
    resp, err := api.client.PostForm(urlGetRouteSummaryForStop, values)
    if err != nil {
        return rsfm, err
    }
    data, err := ioutil.ReadAll(resp.Body)
    if debug {
        fmt.Println(string(data))
    }
    if err != nil {
        return rsfm, err
    }
    rsfm, err = unmarshalRouteSummaryForStop(data)
    if err != nil {
        return rsfm, err
    }
    return rsfm, nil
}

func (api *API) GetNextTripsForStop(stopNo int, routeNo int) (NextTripsForStop, error) {
    var ntfs NextTripsForStop
    values := api.startOfForm()
    values.Add("stopNo", strconv.Itoa(stopNo))
    if routeNo != 0 {
        values.Add("routeNo", strconv.Itoa(routeNo))
    }
    resp, err := api.client.PostForm(urlGetNextTripsForStop, values)
    if err != nil {
        return ntfs, err
    }
    data, err := ioutil.ReadAll(resp.Body)
    if debug {
        fmt.Println(string(data))
    }
    if err != nil {
        return ntfs, err
    }
    ntfs, err = unmarshalNextTripsForStop(data)
    if err != nil {
        return ntfs, err
    }
    return ntfs, nil
}

func (api *API) GetNextTripsForStopAllRoutes(stopNo int) (NextTripsForStopAllRoutes, error) {
    var ntfsas NextTripsForStopAllRoutes
    values := api.startOfForm()
    values.Add("stopNo", strconv.Itoa(stopNo))
    resp, err := api.client.PostForm(urlGetNextTripsForStopAllRoutes, values)
    if err != nil {
        return ntfsas, err
    }
    data, err := ioutil.ReadAll(resp.Body)
    if debug {
        fmt.Println(string(data))
    }
    if err != nil {
        return ntfsas, err
    }
    ntfsas, err = unmarshalNextTripsForStopAllRoutes(data)
    if err != nil {
        return ntfsas, err
    }
    return ntfsas, nil
}

func (api *API) GTFS(table string, id int, column string, value string, order_by string, direction string, limit int) (GTFS, error) {
    var gtfs GTFS
    values := api.startOfForm()
    values.Add("table", table)
    if id != 0 {
        values.Add("id", strconv.Itoa(id))
    }
    if column != "" && value != "" {
        values.Add("column", column)
        values.Add("value", value)
    }
    if order_by != "" {
        values.Add("order_by", order_by)
    }
    if direction != "" {
        values.Add("direction", direction)
    }
    if limit != 0 {
        values.Add("limit", strconv.Itoa(limit))
    }
    resp, err := api.client.PostForm(urlGTFS, values)
    if err != nil {
        return gtfs, err
    }
    data, err := ioutil.ReadAll(resp.Body)
    if debug {
        fmt.Println(string(data))
    }
    if err != nil {
        return gtfs, err
    }
    gtfs, err = unmarshalGTFS(data)
    if err != nil {
        return gtfs, err
    }
    return gtfs, nil
}
