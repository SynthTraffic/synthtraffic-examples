package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestHandlerAcceptsConnectorRequest(t *testing.T) {
	t.Parallel()
	request := httptest.NewRequest(http.MethodPost, "/api/orders/ORD-000001?environment=dev", strings.NewReader(`{"orderId":"ORD-000001"}`))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	newHandler().ServeHTTP(response, request)

	if response.Code != http.StatusAccepted || response.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("response status=%d content-type=%q", response.Code, response.Header().Get("Content-Type"))
	}
	var payload acceptedResponse
	if err := json.Unmarshal(response.Body.Bytes(), &payload); err != nil {
		t.Fatal(err)
	}
	if payload.Status != "accepted" || payload.Method != http.MethodPost || payload.Path != "/api/orders/ORD-000001" || payload.Query.Get("environment") != "dev" || string(payload.Body) != `{"orderId":"ORD-000001"}` {
		t.Fatalf("payload = %#v", payload)
	}
}

func TestHandlerHealthAndSimulatedStatuses(t *testing.T) {
	t.Parallel()
	tests := []struct {
		path string
		want int
	}{
		{path: "/health", want: http.StatusOK},
		{path: "/status/404", want: http.StatusNotFound},
		{path: "/status/500", want: http.StatusInternalServerError},
		{path: "/status/not-a-code", want: http.StatusBadRequest},
	}
	for _, test := range tests {
		request := httptest.NewRequest(http.MethodGet, test.path, nil)
		response := httptest.NewRecorder()
		newHandler().ServeHTTP(response, request)
		if response.Code != test.want {
			t.Fatalf("%s status = %d, want %d", test.path, response.Code, test.want)
		}
	}
}

func TestHandlerRejectsInvalidAndOversizedBodies(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		body string
		want int
	}{
		{name: "invalid JSON", body: "{", want: http.StatusBadRequest},
		{name: "oversized", body: strings.Repeat("x", maxBodyBytes+1), want: http.StatusRequestEntityTooLarge},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/api/orders/test", strings.NewReader(test.body))
			response := httptest.NewRecorder()
			newHandler().ServeHTTP(response, request)
			if response.Code != test.want {
				t.Fatalf("status = %d, want %d", response.Code, test.want)
			}
		})
	}
}

func TestServeStopsCleanlyWhenCanceled(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(t.Context())
	done := make(chan error, 1)
	go func() {
		done <- serve(ctx, "127.0.0.1:0")
	}()
	t.Cleanup(cancel)

	time.Sleep(20 * time.Millisecond)
	cancel()
	select {
	case err := <-done:
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("server did not stop after cancellation")
	}
}

func TestLoggingResponseWriterPreservesStatusAndUnwraps(t *testing.T) {
	t.Parallel()
	recorder := httptest.NewRecorder()
	wrapped := &responseStatusWriter{ResponseWriter: recorder, status: http.StatusOK}
	wrapped.WriteHeader(http.StatusNoContent)
	if wrapped.status != http.StatusNoContent || wrapped.Unwrap() != recorder {
		t.Fatalf("wrapped status=%d underlying=%T", wrapped.status, wrapped.Unwrap())
	}
}

func TestHandlerRejectsUnreadableBody(t *testing.T) {
	t.Parallel()
	request := httptest.NewRequest(http.MethodPost, "/api/orders/test", nil)
	request.Body = failingReadCloser{}
	response := httptest.NewRecorder()
	newHandler().ServeHTTP(response, request)
	if response.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusBadRequest)
	}
}

func TestRunReportsListenFailure(t *testing.T) {
	t.Setenv("HTTP_TEST_SERVER_ADDRESS", "invalid:address:extra")
	if code := run(); code != 1 {
		t.Fatalf("run code = %d, want 1", code)
	}
}

type failingReadCloser struct{}

func (failingReadCloser) Read([]byte) (int, error) { return 0, errors.New("read failed") }
func (failingReadCloser) Close() error             { return nil }

var _ io.ReadCloser = failingReadCloser{}
