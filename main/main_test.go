package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkHTML(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "localhost:9080", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleHTML)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		handler.ServeHTTP(rr, req)
	}
}
