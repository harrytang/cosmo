package trace

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.uber.org/zap"
)

func TestStartAgent(t *testing.T) {

	const (
		endpoint = "localhost:1234"
	)
	c1 := &Config{
		Name: "foo",
	}
	c2 := &Config{
		Name: "bla",
		Exporters: []*Exporter{{
			Endpoint: endpoint,
		}},
	}
	c3 := &Config{
		Name: "otlphttp",
		Exporters: []*Exporter{{
			Endpoint: endpoint,
			Headers: map[string]string{
				"Authorization": "Bearer token",
			},
			HTTPPath: "/v1/traces",
		}},
	}

	log := zap.NewNop()

	_, err := StartAgent(context.Background(), log, c1)
	assert.NoError(t, err)

	_, err = StartAgent(context.Background(), log, c2)
	assert.NoError(t, err)

	_, err = StartAgent(context.Background(), log, c3)
	assert.NoError(t, err)
}
