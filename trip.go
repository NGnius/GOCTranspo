// NGnius 2020-01-23
package goctranspo

import (
    "strconv"
    "time"
)

type tripInterim struct {
    TripDestination string
    TripStartTime string
    AdjustedScheduleTime string
    AdjustmentAge string
    LastTripOfSchedule bool
    BusType string
    Latitude string
    Longitude string
    GPSSpeed string
}

func (t tripInterim) Fix() Trip {
    tripStartTime, _ := time.Parse("15:04", t.TripStartTime) // HH:mm (HH in 24-hour format)
    adjustedScheduleTime, _ := strconv.Atoi(t.AdjustedScheduleTime)
    adjustmentAge, _ := strconv.ParseFloat(t.AdjustmentAge, 32)
    latitude, _ := strconv.ParseFloat(t.Latitude, 64)
    longitude, _ := strconv.ParseFloat(t.Longitude, 64)
    gpsSpeed, _ := strconv.ParseFloat(t.GPSSpeed, 32)
    return Trip {
                TripDestination: t.TripDestination,
                TripStartTime: tripStartTime,
                AdjustedScheduleTime: adjustedScheduleTime,
                AdjustmentAge: adjustmentAge,
                LastTripOfSchedule: t.LastTripOfSchedule,
                BusType: t.BusType,
                Latitude: latitude,
                Longitude: longitude,
                GPSSpeed: gpsSpeed,
                }
}

type Trip struct {
    TripDestination string
    TripStartTime time.Time
    AdjustedScheduleTime int
    AdjustmentAge float64
    LastTripOfSchedule bool
    BusType string
    Latitude float64
    Longitude float64
    GPSSpeed float64
}
