package handler

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_getAllTags(t *testing.T) {
	tests := []struct {
		name                 string
		mockBehavior         func(r *mock_service.MockTag)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Success",
			mockBehavior: func(r *mock_service.MockTag) {
				tags := []string{"tag1", "tag2", "tag3"}
				r.EXPECT().GetAllTags(gomock.Any()).Return(tags, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":200,"message":"success","payload":["tag1","tag2","tag3"]}`,
		},
		{
			name: "Internal Server Error",
			mockBehavior: func(r *mock_service.MockTag) {
				r.EXPECT().GetAllTags(gomock.Any()).Return(nil, errors.New("Internal Server Error"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"status":500,"message":"Internal Server Error","payload":""}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			tagService := mock_service.NewMockTag(c)
			test.mockBehavior(tagService)

			ctx := context.Background()
			services := &service.Service{Tag: tagService}
			handler := Handler{services, &ctx}

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/api/v1/tags", nil)

			handler.getAllTags(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
