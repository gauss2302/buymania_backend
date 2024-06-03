package user_service

import (
	"bytes"
	"encoding/json"
	"github.com/gauss2302/buymania_backend/types"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("come with an error if user id is not a number", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/user/xyz", nil)
		if err != nil {
			t.Fatal(err)
		}

		record := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/user/{userID}", handler.handleGetUser).Methods(http.MethodGet)

		router.ServeHTTP(record, req)

		if record.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, record.Code)
		}

	})

	t.Run("should handle users by Id", func(t *testing.T) {

	})

	t.Run("should fails is the user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "user",
			LastName:  "123",
			Email:     "nikita@gmail.com",
			Password:  "asd",
		}

		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		record := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(record, req)

		if record.Code != http.StatusBadRequest {
			t.Errorf("expectd status code %d, got %d", http.StatusBadRequest, record.Code)
		}
	})
}

type mockUserStore struct {
}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return &types.User{}, nil
}

func (m *mockUserStore) CreateUser(u types.User) error {
	return nil
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return &types.User{}, nil
}
