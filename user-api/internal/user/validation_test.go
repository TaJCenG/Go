package user

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestValidate(t *testing.T) {
	u := User{Name: "Taj", Email: "taj@example.com"}
	if err := Validate(u); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestValidate_TableDriven(t *testing.T) {
	tests := []struct {
		name    string
		input   User
		wantErr bool
	}{
		{"valid user", User{Name: "Taj", Email: "taj@example.com"}, false},
		{"missing name", User{Name: "", Email: "taj@example.com"}, true},
		{"missing email", User{Name: "Taj", Email: ""}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Validate(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("got error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type MockRepo struct {
	users map[int]User
}

func (m *MockRepo) GetByID(id int) (User, error) {
	u, ok := m.users[id]
	if !ok {
		return User{}, errors.New("not found")
	}
	return u, nil
}

func TestService_GetUser(t *testing.T) {
	mock := &MockRepo{users: map[int]User{1: {ID: 1, Name: "Taj", Email: "taj@example.com"}}}
	service := NewService(mock)

	u, err := service.GetByID(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if u.Name != "Taj" {
		t.Errorf("expected Taj, got %s", u.Name)
	}
}
func TestGetUserHandler(t *testing.T) {
	mock := &MockRepo{users: map[int]User{1: {ID: 1, Name: "Taj", Email: "taj@example.com"}}}
	service := NewService(mock)
	handler := NewHandler(service)

	req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	w := httptest.NewRecorder()

	handler.UserByID(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}
