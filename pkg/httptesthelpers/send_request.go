package httptesthelpers

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func SendRequest(method string, url string, body io.Reader, router *gin.Engine, headers map[string]string) (int, *bytes.Buffer, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return 0, nil, err
	}

	for header, value := range headers {
		req.Header.Set(header, value)
	}

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	return rec.Code, rec.Body, nil
}
