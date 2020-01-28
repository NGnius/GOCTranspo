// NGnius 2020-01-24

package goctranspo

import (
    "strconv"
)

type queryInterim struct {
    Table string `json:"table"`
    Order_By string `json:"order_by"`
    Direction string `json:"direction"`
    Column string `json:"column"`
    Value string `json:"value"`
    Limit string `json:"limit"`
    Format string `json:"format"`
}

func (q queryInterim) Fix() Query {
    limit, _ := strconv.Atoi(q.Limit)
    return Query {
                Table: q.Table,
                Order_By: q.Order_By,
                Direction: q.Direction,
                Column: q.Column,
                Value: q.Value,
                Limit: limit,
                Format: q.Format,
                }
}

type Query struct {
    Table string
    Order_By string
    Direction string
    Column string
    Value string
    Limit int
    Format string
}
