// Command http-server runs a small local destination for the HTTP connector example.
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

const (
	defaultAddress = ":18081"
	maxBodyBytes   = 1 << 20
)

func main() {
	os.Exit(run())
}

func run() int {
	address := os.Getenv("HTTP_TEST_SERVER_ADDRESS")
	if address == "" {
		address = defaultAddress
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	if err := serve(ctx, address); err != nil {
		slog.Error("HTTP test server failed", "error", err)
		return 1
	}
	return 0
}

func serve(ctx context.Context, address string) error {
	server := &http.Server{
		Addr:              address,
		Handler:           logRequests(newHandler()),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       30 * time.Second,
	}
	shutdownDone := make(chan error, 1)
	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		shutdownDone <- server.Shutdown(shutdownCtx)
	}()

	slog.Info("HTTP test server listening", "address", address)
	err := server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	if shutdownErr := <-shutdownDone; shutdownErr != nil {
		return fmt.Errorf("shut down HTTP test server: %w", shutdownErr)
	}
	return nil
}

func newHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handleHealth)
	mux.HandleFunc("/status/", handleStatus)
	mux.HandleFunc("/api/orders/", handleRequest)
	return mux
}

func handleHealth(writer http.ResponseWriter, _ *http.Request) {
	writeJSON(writer, http.StatusOK, map[string]any{"status": "ok"})
}

func handleStatus(writer http.ResponseWriter, request *http.Request) {
	value := strings.TrimPrefix(request.URL.Path, "/status/")
	status, err := strconv.Atoi(value)
	if err != nil || status < 200 || status > 599 {
		writeJSON(writer, http.StatusBadRequest, map[string]any{
			"status":  "rejected",
			"message": "status must be an integer from 200 through 599",
		})
		return
	}
	writeJSON(writer, status, map[string]any{
		"status":     "simulated",
		"httpStatus": status,
	})
}

func handleRequest(writer http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(http.MaxBytesReader(writer, request.Body, maxBodyBytes))
	if err != nil {
		var tooLarge *http.MaxBytesError
		if errors.As(err, &tooLarge) {
			writeJSON(writer, http.StatusRequestEntityTooLarge, map[string]any{"status": "rejected", "message": "request body exceeds 1 MiB"})
			return
		}
		writeJSON(writer, http.StatusBadRequest, map[string]any{"status": "rejected", "message": "could not read request body"})
		return
	}

	trimmedBody := strings.TrimSpace(string(body))
	if trimmedBody != "" && !json.Valid(body) {
		writeJSON(writer, http.StatusBadRequest, map[string]any{"status": "rejected", "message": "request body must be valid JSON"})
		return
	}

	response := acceptedResponse{
		Status:  "accepted",
		Message: "request received",
		Method:  request.Method,
		Path:    request.URL.Path,
		Query:   request.URL.Query(),
	}
	if trimmedBody != "" {
		response.Body = json.RawMessage(body)
	}
	writeJSON(writer, http.StatusAccepted, response)
}

type acceptedResponse struct {
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Method  string          `json:"method"`
	Path    string          `json:"path"`
	Query   url.Values      `json:"query,omitempty"`
	Body    json.RawMessage `json:"body,omitempty"`
}

func writeJSON(writer http.ResponseWriter, status int, value any) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	if err := json.NewEncoder(writer).Encode(value); err != nil {
		slog.Error("HTTP test server response failed", "error", err)
	}
}

func logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		statusWriter := &responseStatusWriter{ResponseWriter: writer, status: http.StatusOK}
		next.ServeHTTP(statusWriter, request)
		slog.Info("HTTP test request", "method", request.Method, "path", request.URL.Path, "status", statusWriter.status)
	})
}

type responseStatusWriter struct {
	http.ResponseWriter
	status int
}

func (writer *responseStatusWriter) WriteHeader(status int) {
	writer.status = status
	writer.ResponseWriter.WriteHeader(status)
}

func (writer *responseStatusWriter) Unwrap() http.ResponseWriter {
	return writer.ResponseWriter
}
