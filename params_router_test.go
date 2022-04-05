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

func TestRouterParams(t *testing.T) {

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

func TestRouterIdParams(t *testing.T) {

	router := httprouter.New()
	router.GET("/products/:id/items/:itemId", func(w http.ResponseWriter, request *http.Request, p httprouter.Params) {
		id := p.ByName("id")
		itemId := p.ByName("itemId")
		text := "Product " + id + " Item " + itemId
		fmt.Fprint(w, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:9090/products/1/items/1", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Product 1 Item 1", string(body))
}

func TestRouterIdParamsCacthAll(t *testing.T) {

	router := httprouter.New()
	router.GET("/images/*image", func(w http.ResponseWriter, request *http.Request, p httprouter.Params) {
		image := p.ByName("image")
		text := "Image: " + image
		fmt.Fprint(w, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:9090/images/small/local.jpeg", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Image: /small/local.jpeg", string(body))
}
