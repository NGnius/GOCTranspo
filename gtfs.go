// NGnius 2020-01-24

package goctranspo

import (
    "encoding/json"
)

type gtfsInterim struct {
    Query queryInterim
    Gtfs []map[string]string
}

func (g gtfsInterim) Fix() GTFS {
    return GTFS {
                Query: g.Query.Fix(),
                Gtfs: g.Gtfs,
                }
}

type GTFS struct {
    Query Query
    Gtfs []map[string]string
}

func unmarshalGTFS(data []byte) (GTFS, error) {
    var pseudoGtfs gtfsInterim
    var gtfs GTFS
    err := json.Unmarshal(data, &pseudoGtfs)
    if err != nil {
        return gtfs, err
    }
    gtfs = pseudoGtfs.Fix()
    return gtfs, err
}
