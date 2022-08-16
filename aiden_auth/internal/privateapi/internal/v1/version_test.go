package v1_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeReq(t *testing.T, method, path string, body interface{}) (*httptest.ResponseRecorder, *http.Request) {
	var (
		bs  []byte
		err error
	)
	if body != nil {
		bs, err = json.Marshal(body)
		assert.NoError(t, err)
	}

	req, err := http.NewRequest(method, "http://test/"+path, bytes.NewReader(bs))
	assert.NoError(t, err)
	return httptest.NewRecorder(), req
}
