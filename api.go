// NGnius 2020-01-23

package goctranspo

import (
    "fmt"
    "net/http"
    "net/url"
    "strconv"
    "io/ioutil"
    
    "errors"
)

const (
    urlGetRouteSummaryForStop = "https://api.octranspo1.com/v1.3/GetRouteSummaryForStop"
    urlGetNextTripsForStop = "https://api.octranspo1.com/v1.3/GetNextTripsForStop"
    urlGetNextTripsForStopAllRoutes = "https://api.octranspo1.com/v1.3/GetNextTripsForStopAllRoutes"
    urlGTFS = "https://api.octranspo1.com/v1.3/Gtfs"
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
    //fmt.Println(string(data))
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
    //fmt.Println(string(data))
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
    //fmt.Println(string(data))
    if err != nil {
        return ntfsas, err
    }
    ntfsas, err = unmarshalNextTripsForStopAllRoutes(data)
    if err != nil {
        return ntfsas, err
    }
    return ntfsas, nil
}

func (api *API) GTFS() (error) {
    fmt.Println("GTFS query support coming soon")
    return errors.New("UNIMPLEMENTED")
}
