package main

import (
	"github.com/influxdata/influxdb/client/v2"
)

const (
	tableName = "boat"

	dbName = "telemetry"

	addr = "http://influxdb:8086"
)

func save(points [][]interface{}) error {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: addr,
	})
	if err != nil {
		return err
	}
	defer c.Close()

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  dbName,
		Precision: "s",
	})
	if err != nil {
		return err
	}

	for _, point := range points {
		e := newEvent(point)

		fields := e.toFields()
		tags := map[string]string{}

		ts, err := e.Time()
		if err != nil {
			return err
		}

		pt, err := client.NewPoint(tableName, tags, fields, ts)
		if err != nil {
			return err
		}
		bp.AddPoint(pt)
	}

	return c.Write(bp)
}
