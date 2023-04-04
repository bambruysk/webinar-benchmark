package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkJSON(b *testing.B) {
	reqBody := []byte(`{"name": "John", "age": 30}`)
	req, _ := http.NewRequest(http.MethodPost, "localhost:9080", bytes.NewBuffer(reqBody))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleJSON)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		handler.ServeHTTP(rr, req)
	}
}
