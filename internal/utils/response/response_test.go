package response_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TheSaifHub/Student-Records-REST-API/internal/utils/response"
)

func TestWriteJson(t *testing.T) {

	rec := httptest.NewRecorder()

	err := response.WriteJson(rec, 200, map[string]string{
		"status": "ok",
	})

	if err != nil {
		t.Fatal("expected no error, got:", err)
	}

	if rec.Code != 200 {
		t.Errorf("expected status 200, got %d", rec.Code)
	}
}

func TestGeneralError(t *testing.T) {

	errResp := response.GeneralError(
		http.ErrAbortHandler,
	)

	if errResp.Status != "Error" {
		t.Errorf("expected Error status, got %s", errResp.Status)
	}
}
