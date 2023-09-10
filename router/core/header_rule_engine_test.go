package core

import (
	"github.com/stretchr/testify/assert"
	"github.com/wundergraph/cosmo/router/internal/config"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNamedPropagateHeaderRule(t *testing.T) {

	ht := NewHeaderTransformer(config.HeaderRules{
		All: config.GlobalHeaderRule{
			Request: []config.RequestHeaderRule{
				{
					Operation: "propagate",
					Named:     "X-Test-1",
				},
			},
		},
	})

	rr := httptest.NewRecorder()

	clientReq, err := http.NewRequest("POST", "http://localhost", nil)
	clientReq.Header.Set("X-Test-1", "test1")
	clientReq.Header.Set("X-Test-2", "test2")

	originReq, err := http.NewRequest("POST", "http://localhost", nil)
	assert.Nil(t, err)

	updatedClientReq, _ := ht.OnOriginRequest(originReq, &requestContext{
		logger:         zap.NewNop(),
		responseWriter: rr,
		request:        clientReq,
		operation:      &operationContext{},
	})

	assert.Equal(t, "test1", updatedClientReq.Header.Get("X-Test-1"))
	assert.Empty(t, updatedClientReq.Header.Get("X-Test-2"))
}

func TestRegexPropagateHeaderRule(t *testing.T) {

	ht := NewHeaderTransformer(config.HeaderRules{
		All: config.GlobalHeaderRule{
			Request: []config.RequestHeaderRule{
				{
					Operation: "propagate",
					Matching:  "(?i)X-Test-.*",
				},
			},
		},
	})

	rr := httptest.NewRecorder()

	clientReq, err := http.NewRequest("POST", "http://localhost", nil)
	clientReq.Header.Set("X-Test-1", "test1")
	clientReq.Header.Set("X-Test-2", "test2")
	clientReq.Header.Set("Y-Test", "test3")

	originReq, err := http.NewRequest("POST", "http://localhost", nil)
	assert.Nil(t, err)

	updatedClientReq, _ := ht.OnOriginRequest(originReq, &requestContext{
		logger:         zap.NewNop(),
		responseWriter: rr,
		request:        clientReq,
		operation:      &operationContext{},
	})

	assert.Equal(t, "test1", updatedClientReq.Header.Get("X-Test-1"))
	assert.Equal(t, "test2", updatedClientReq.Header.Get("X-Test-2"))
	assert.Empty(t, updatedClientReq.Header.Get("Y-Test"))
}