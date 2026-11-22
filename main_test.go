package main

import (
	"encoding/json/v2"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetAccountData(t *testing.T) {
	router := newRouter()

	for _, accountID := range []string{"test_account_id", "account-123", "42"} {
		t.Run(accountID, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/%s/data", accountID), nil)
			rec := httptest.NewRecorder()

			start := time.Now().Unix()
			router.ServeHTTP(rec, req)
			end := time.Now().Unix()

			if rec.Code != http.StatusOK {
				t.Fatalf("status = %d, want %d, body = %s", rec.Code, http.StatusOK, rec.Body.String())
			}

			var got AccountData
			if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
				t.Fatal(err)
			}
			if got.AccountID != accountID {
				t.Errorf("accountId = %q, want %q", got.AccountID, accountID)
			}
			if got.Timestamp < start || got.Timestamp > end {
				t.Errorf("timestamp = %d, want between %d and %d", got.Timestamp, start, end)
			}
			if want := "Celtra DevOps"; got.Data != want {
				t.Errorf("data = %q, want %q", got.Data, want)
			}
		})
	}
}

func TestInvalidAccountDataRequests(t *testing.T) {
	router := newRouter()
	requests := []struct {
		name   string
		method string
		path   string
	}{
		{name: "missing_account_id", method: http.MethodGet, path: "/data"},
		{name: "unsupported_method", method: http.MethodPost, path: "/account-123/data"},
	}

	for _, request := range requests {
		t.Run(request.name, func(t *testing.T) {
			req := httptest.NewRequest(request.method, request.path, nil)
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)

			if rec.Code != http.StatusNotFound {
				t.Errorf("status = %d, want %d, body = %s", rec.Code, http.StatusNotFound, rec.Body.String())
			}
		})
	}
}

func TestReady(t *testing.T) {
	router := newRouter()
	req := httptest.NewRequest(http.MethodGet, "/ready", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d, body = %s", rec.Code, http.StatusOK, rec.Body.String())
	}
	if want := "Ready"; rec.Body.String() != want {
		t.Errorf("body = %q, want %q", rec.Body.String(), want)
	}
}
