package main

import "google.golang.org/protobuf/proto"

type Response struct {
	Data *ResponseData `json:"data"`
}

type ResponseData struct {
	Metrics []*Metric `json:"metrics"`
	Stats   Stats     `json:"stats"`
}

type Metric struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type Stats struct {
	Sales *int64 `json:"sales,omitempty"` //total sales
}

var (
	example = Response{
		Data: &ResponseData{
			Metrics: []*Metric{
				{
					Name:  "gmv",
					Value: 123,
				},
			},
			Stats: Stats{
				Sales: proto.Int64(100),
			},
		},
	}
)
