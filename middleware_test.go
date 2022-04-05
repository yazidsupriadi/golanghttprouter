package main_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

type LogMiddleware struct {
	http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request Accepted")
	middleware.Handler.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T) {

	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, request *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Middleware")
	})

	middleware := LogMiddleware{router}
	request := httptest.NewRequest("GET", "http://localhost:9090/", nil)
	recorder := httptest.NewRecorder()
	middleware.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Middleware", string(body))
}
