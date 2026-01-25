package student_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TheSaifHub/Student-Records-REST-API/internal/http/handlers/student"
	"github.com/TheSaifHub/Student-Records-REST-API/internal/types"
)

type MockStorage struct{}

func (m *MockStorage) CreateStudent(name string, email string, age int) (int64, error) {
	return 1, nil
}

func (m *MockStorage) GetStudentById(id int64) (types.Student, error) {
	return types.Student{
		Id:    id,
		Name:  "Saif",
		Email: "saif@gmail.com",
		Age:   22,
	}, nil
}

func (m *MockStorage) GetStudents() ([]types.Student, error) {
	return []types.Student{
		{
			Id:    1,
			Name:  "Saif",
			Email: "saif@gmail.com",
			Age:   22,
		},
		{
			Id:    2,
			Name:  "Ali",
			Email: "ali@gmail.com",
			Age:   21,
		},
	}, nil
}


func TestCreateStudent_Success(t *testing.T) {
	storage := &MockStorage{}

	handler := student.New(storage)

	body := []byte(`{
		"name": "Saif",
		"email": "saif@gmail.com",
		"age": 22
	}`)

	req := httptest.NewRequest(http.MethodPost, "/api/students", bytes.NewBuffer(body))
	rec := httptest.NewRecorder()

	handler(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected status 201, got %d", rec.Code)
	}
}

func TestCreateStudent_EmptyBody(t *testing.T) {
	storage := &MockStorage{}
	handler := student.New(storage)

	req := httptest.NewRequest(http.MethodPost, "/api/students", nil)
	rec := httptest.NewRecorder()

	handler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", rec.Code)
	}
}

func TestGetStudentList_Success(t *testing.T) {
	storage := &MockStorage{}
	handler := student.GetList(storage)

	req := httptest.NewRequest(http.MethodGet, "/api/students/list", nil)
	rec := httptest.NewRecorder()

	handler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}
}
