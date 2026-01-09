package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEchoHandler(t *testing.T) {
	// Body requestu, które chcemy wysłać
	body := "Hello Echo Server!"
	req := httptest.NewRequest("POST", "/", strings.NewReader(body))
	w := httptest.NewRecorder() // fake ResponseWriter

	// Wywołanie handlera
	echoHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	// Sprawdzenie status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	// Odczytanie body odpowiedzi
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("cannot read response body: %v", err)
	}

	// Porównanie z tym, co wysłaliśmy
	if string(respBody) != body {
		t.Errorf("expected body %q, got %q", body, string(respBody))
	}
}
