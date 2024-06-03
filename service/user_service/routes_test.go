package user_service

import (
	"net/http"
	"testing"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := *mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("come with an error if user id is not a number", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/user/xyz", nil)

		if err != nil {
			t.Fatal(err)
		}
	})
}

type mockUserStore struct {
}
