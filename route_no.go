// NGnius 2020-01-24

package goctranspo

import (
    "strconv"
)

type RouteNo struct {
    Number int
    String string
    IsAlphanumeric bool
}

func NewRouteNo(route string) RouteNo {
    newRN := RouteNo {}
    num, err := strconv.Atoi(route)
    if err != nil {
        newRN.IsAlphanumeric = true
    } else {
        newRN.Number = num
        newRN.IsAlphanumeric = false
    }
    newRN.String = route
    return newRN
}
