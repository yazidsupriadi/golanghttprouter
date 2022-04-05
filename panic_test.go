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

func TestPanic(t *testing.T) {

	router := httprouter.New()

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, err interface{}) {
		fmt.Fprint(w, "Panic : ", err)
	}

	router.GET("/", func(w http.ResponseWriter, request *http.Request, p httprouter.Params) {
		panic("UPS")
	})

	request := httptest.NewRequest("GET", "http://localhost:9090/", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Panic : UPS", string(body))
}
