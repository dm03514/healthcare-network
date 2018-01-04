package main

import (
	"net/http"
	"testing"
	"fmt"
	"encoding/json"
)

func TestVisJSHealthNetwork_NodesEdges(t *testing.T) {
	hn := NewHealthNetwork(map[string]Health{
		"E": {
			Latency:    54,
			Service:    "E",
			StatusCode: http.StatusOK,
			Integrations: NewHealthNetwork(map[string]Health{
				"cassandra": {
					Latency:    10,
					Service:    "cassandra",
					StatusCode: http.StatusOK,
				},
				"ew": {
					Latency:    10,
					Service:    "ew",
					StatusCode: http.StatusOK,
					Integrations: NewHealthNetwork(map[string]Health{
						"s3": {
							Latency:    10,
							Service:    "s3",
							StatusCode: http.StatusOK,
						},
					}),
				},
			}),
		},
		"vw": {
			Latency:    54,
			Service:    "vw",
			StatusCode: http.StatusOK,
			Integrations: NewHealthNetwork(map[string]Health{
				"s3": {
					Latency:    10,
					Service:    "s3",
					StatusCode: http.StatusOK,
				},
				"vertica": {
					Latency:    10,
					Service:    "vertica",
					StatusCode: http.StatusOK,
				},
			}),
		},
	})
	vis := VisJSGraph{Graph: hn.Graph(nil)}
	fmt.Printf("%+v\n", vis.Nodes())
	s, err := json.Marshal(vis.Nodes())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", s)

	es, err := json.Marshal(vis.Edges())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", es)
}
