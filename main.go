package main

import (
	"fmt"
)

/*
{
	service: "service 1",
	latency: X_ms
	status_code: 200OK,

	integrations: {
		(nsqd): {
			status_code: 200OK,
			latency: X_ms
		}
	}
}


{
	status_code: 200OK,
	service: "Service 2",

	integrations: {
		(nsqd): {
			status_code: 200OK,
			latency: X_ms
		},
		(vertica): {
			status_code: 200OK,
			latency: X_ms
		}
		(s3): {
			status_code: 200OK,
			latency: X_ms
		}
	}
}


{
	status_code: 200OK,
	service: "service 3"

	integrations: {
		service 4: {
			status_code: 200OK,
			latency: X_ms,

			integrations: {
				mongo: {
					status_code: 200OK,
					latency: X_ms
				}
			}
		}
	}
}
*/


type Graph map[string][]string

func (g Graph) Nodes() []string {
	ns := []string{}
	for s := range g {
		ns = append(ns, s)
	}
	return ns
}


func NewHealthNetwork(ServiceHealths map[string]Health) *HealthNetwork {
	return &HealthNetwork{
		ServiceHealths: ServiceHealths,
	}
}

type HealthNetwork struct {
	ServiceHealths map[string]Health
	graph Graph
}

func (hn HealthNetwork) Services() []string {
	svcs := make([]string, 0, len(hn.ServiceHealths))
	for s := range hn.ServiceHealths {
		svcs = append(svcs, s)
	}
	return svcs
}

func (hn *HealthNetwork) Graph(graph Graph) Graph {
	// if client did not pass in a graph initialize a new one
	fmt.Printf("Graph is %+v\n", graph)
	if graph == nil {
		hn.graph = make(map[string][]string)
	} else {
		hn.graph = graph
	}

	// recurse subgraphs until we've
	for _, h := range hn.ServiceHealths {
		// base case integration is nil, done with this service
		if h.Integrations == nil {
			continue
		}

		// if not initialize and add integration
		for _, s := range h.Integrations.Services() {
			hn.graph[s] = []string{}
			hn.graph[h.Service] = append(hn.graph[h.Service], s)
		}

		// recurse through all integrations
		h.Integrations.Graph(hn.graph)
	}
	return hn.graph
}

type Health struct {
	Latency      int
	Service      string
	StatusCode   int
	Integrations *HealthNetwork
}

func main() {
	fmt.Println("Hello")
}
