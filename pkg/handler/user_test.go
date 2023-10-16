package handler

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_user(t *testing.T) {
	mockCookie := &http.Cookie{
		Name:  "session_id",
		Value: "value",
	}
	mockUser := model.User{Mail: "max@max.ru", PasswordHash: "passWord", Name: "Max"}
	jsonData, _ := json.Marshal(mockUser)

	tests := []struct {
		name                 string
		inputCookie          *http.Cookie
		mockBehavior         func(r *mock_service.MockUser)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "Ok",
			inputCookie: mockCookie,
			mockBehavior: func(r *mock_service.MockUser) {
				r.EXPECT().GetCurrentUser(gomock.Any(), mockCookie.Value).Return(mockUser, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: string(jsonData),
		},
		{
			name:                 "No cookie",
			inputCookie:          nil,
			mockBehavior:         func(r *mock_service.MockUser) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"no session"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockUser(c)
			test.mockBehavior(repo)

			services := &service.Service{User: repo}
			handler := Handler{services}

			mux := http.NewServeMux()
			mux.HandleFunc("/api/user", handler.user)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/api/user", nil)
			if test.inputCookie != nil {
				req.AddCookie(test.inputCookie)
			}

			mux.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}