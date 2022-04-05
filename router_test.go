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

func TestRouter(t *testing.T) {

	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, request *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "hello")
	})

	request := httptest.NewRequest("GET", "http://localhost:9090/", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "hello", string(body))
}

func TestRouterId(t *testing.T) {

	router := httprouter.New()
	router.GET("/products/:id", func(w http.ResponseWriter, request *http.Request, p httprouter.Params) {
		id := p.ByName("id")
		text := "Product " + id
		fmt.Fprint(w, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:9090/products/1", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Product 1", string(body))
}
