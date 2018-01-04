package main

import (
	"testing"
	"net/http"
	"github.com/stretchr/testify/assert"
)

func Test_Health_JSONDecodesNoIntegrations(t *testing.T) {
	t.Skip()
}

func Test_Health_JSONDecodes_SingleIntegration(t *testing.T) {
	t.Skip()
}

func Test_Health_JSONDecodes_MultipleIntegrations(t *testing.T) {
	t.Skip()
}

func Test_Health_JSONDecodes_SingleNestedIntegrations(t *testing.T) {
	t.Skip()
}

func TestHealthNetwork_Graph(t *testing.T) {
	hn := NewHealthNetwork(map[string]Health{
		"X": {
			Latency:    54,
			Service:    "X",
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

	assert.Equal(t,
		Graph(
			map[string][]string{
				"X": []string{"s3", "vertica"},
			},
		),
		hn.Graph(nil),
	)
}
