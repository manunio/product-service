package handlers

import (
	"bytes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"product-service/data"
	"testing"
)

var r *mux.Router
var l *log.Logger
var v *data.Validation
var rr *httptest.ResponseRecorder
var ph *Products

// Integration tests

func setup() {
	l = log.New(os.Stdout, "product-service", log.LstdFlags)
	v = data.NewValidation()

	// create the handlers
	ph = NewProducts(l, v)
	rr = httptest.NewRecorder()

	r = mux.NewRouter()
}

func TestGetListAll(t *testing.T) {
	setup()
	req, _ := http.NewRequest(http.MethodGet, "/products", nil)
	r.HandleFunc("/products", ph.ListAll).Methods(http.MethodGet)
	r.ServeHTTP(rr, req)
	checkResponseCode(t, http.StatusOK, rr.Code)
}

func TestGetListSingle(t *testing.T) {
	setup()
	req, _ := http.NewRequest(http.MethodGet, "/products/1", nil)
	r.HandleFunc("/products/{id:[0-9]+}", ph.ListSingle).Methods(http.MethodGet)
	r.ServeHTTP(rr, req)
	checkResponseCode(t, http.StatusOK, rr.Code)
}

func TestPost(t *testing.T) {
	setup()
	jsonStr := []byte(`{"name":"tea","description":"a cup of tea","price":10,"sku":"abc-abc-abc"}`)
	req, _ := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	r.HandleFunc("/products", ph.Create).Methods(http.MethodPost)
	r.Use(ph.MiddlewareValidateProduct)
	r.ServeHTTP(rr, req)
	checkResponseCode(t, http.StatusOK, rr.Code)
}

func TestPut(t *testing.T) {
	setup()
	jsonStr := []byte(`{"id":2,"name":"tea","description":"a cup of tea","price":10,"sku":"abc-abc-abc"}`)
	req, _ := http.NewRequest(http.MethodPut, "/products", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	r.HandleFunc("/products", ph.Update).Methods(http.MethodPut)
	r.Use(ph.MiddlewareValidateProduct)
	r.ServeHTTP(rr, req)
	checkResponseCode(t, http.StatusNoContent, rr.Code)
}

func TestDelete(t *testing.T) {
	setup()
	req, _ := http.NewRequest(http.MethodDelete, "/products/1", nil)
	r.HandleFunc("/products/{id:[0-9]+}", ph.Delete).Methods(http.MethodDelete)
	r.ServeHTTP(rr, req)
	checkResponseCode(t, http.StatusNoContent, rr.Code)
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {

	l := log.New(os.Stdout, "product-service", log.LstdFlags)
	v := data.NewValidation()

	// create the handlers
	ph := NewProducts(l, v)
	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/products", ph.ListAll).Methods(http.MethodGet)
	r.ServeHTTP(rr, req)
	return rr
}
